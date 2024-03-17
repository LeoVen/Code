use std::{
    net::{Ipv4Addr, SocketAddrV4},
    sync::Arc,
};

use anyhow::{Context, Result};
use axum::{
    body::Body,
    extract::{MatchedPath, Path, State},
    http::{self, Request},
    response::Response,
    Json, Router,
};
use redis::{aio::MultiplexedConnection, AsyncCommands, SetOptions};
use redis_macros::{FromRedisValue, ToRedisArgs};
use serde::{Deserialize, Serialize};
use serde_json::{json, Value};
use sqlx::{FromRow, Postgres};
use tower_http::{cors::CorsLayer, trace::TraceLayer};

use crate::{
    error::{ApiError, ApiResult},
    metrics,
};

#[derive(Serialize, Deserialize, Debug)]
struct AxumApiConfig {
    #[serde(rename(deserialize = "api_axum_port"))]
    pub port: u16,
}

#[derive(Clone)]
pub struct AppState {
    pub db: sqlx::Pool<Postgres>,
    pub cache: MultiplexedConnection,
}

#[tracing::instrument(skip(state))]
pub async fn setup(env: &str, state: AppState) -> Result<()> {
    let config = envy::from_env::<AxumApiConfig>().context("Failed to get env vars")?;

    if env == "dev" {
        let config_str = serde_json::to_string(&config).unwrap_or("Serialize Error".to_string());
        tracing::info!(config = config_str);
    }

    let app = Router::new()
        .route("/", axum::routing::get(select_all))
        .route("/:id", axum::routing::get(select_one))
        .layer(CorsLayer::new().allow_headers([http::header::CONTENT_TYPE]))
        .layer(
            TraceLayer::new_for_http()
                .make_span_with(|request: &Request<_>| {
                    let matched_path = request
                        .extensions()
                        .get::<MatchedPath>()
                        .map(MatchedPath::as_str);

                    tracing::info_span!(
                        "http_request",
                        method = ?request.method(),
                        matched_path,
                    )
                })
                .on_response(
                    |response: &Response<Body>, _duration, _span: &tracing::Span| {
                        let status = response.status();
                        if status.is_success() {
                            tracing::info!(metric = metrics::RESPONSE_STATUS_2XX);
                        } else if status.is_client_error() {
                            tracing::info!(metric = metrics::RESPONSE_STATUS_4XX);
                        } else if status.is_server_error() {
                            tracing::info!(metric = metrics::RESPONSE_STATUS_5XX);
                        }
                    },
                ),
        )
        .with_state(state.into());

    let listener =
        tokio::net::TcpListener::bind(SocketAddrV4::new(Ipv4Addr::new(0, 0, 0, 0), config.port))
            .await
            .context("Failed to bind to TCP port")?;

    tracing::info!("Listening on port {}", config.port);

    axum::serve(listener, app)
        .await
        .context("Axum serve failed")?;

    Ok(())
}

#[derive(Serialize, Deserialize, FromRow, Clone, FromRedisValue, ToRedisArgs)]
pub struct Spell {
    pub id: i64,
    pub name: Option<String>,
    pub damage: i32,
    pub created_at: chrono::DateTime<chrono::Utc>,
    pub updated_at: chrono::DateTime<chrono::Utc>,
}

impl Spell {
    pub fn redis_key(id: i64) -> String {
        format!("{}:{}", "spell", id)
    }
}

#[tracing::instrument(skip_all)]
async fn select_all(State(ctx): State<Arc<AppState>>) -> ApiResult<Json<Value>> {
    let result: Vec<Spell> =
        sqlx::query_as("SELECT id, name, damage, created_at, updated_at FROM spell")
            .fetch_all(&ctx.db)
            .await?;

    Ok(Json(json!(result)))
}

#[tracing::instrument(skip_all)]
async fn select_one(
    State(ctx): State<Arc<AppState>>,
    Path(id): Path<i64>,
) -> ApiResult<Json<String>> {
    tracing::info!(metric = metrics::API_CALL);

    let mut cache = ctx.cache.clone();

    let cached: Option<Vec<String>> = cache.get(Spell::redis_key(id)).await?;

    if let Some(mut cached) = cached {
        tracing::info!(metric = metrics::CACHE_HIT);
        return Ok(Json(cached.pop().unwrap_or("{}".to_string())));
    } else {
        tracing::info!(metric = metrics::CACHE_MISS);
    }

    let value: Spell =
        sqlx::query_as("SELECT id, name, damage, created_at, updated_at FROM spell where id = $1")
            .bind(id)
            .fetch_one(&ctx.db)
            .await
            .map_err(|err| match err {
                sqlx::Error::RowNotFound => ApiError::NotFound(err.to_string()),
                _ => ApiError::Internal(err.into()),
            })?;

    {
        let value = value.clone();
        let mut cache = cache.clone();

        tokio::spawn(async move {
            let opts = SetOptions::default().with_expiration(redis::SetExpiry::EX(60));

            if let Ok(data) = serde_json::to_string(&value) {
                let _: Result<(), _> = cache.set_options(Spell::redis_key(id), data, opts).await;
            }
        });
    }

    Ok(Json(serde_json::to_string(&value).unwrap_or_default()))
}
