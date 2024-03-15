use std::{
    net::{Ipv4Addr, SocketAddrV4},
    sync::Arc,
};

use anyhow::{Context, Result};
use axum::{
    extract::{MatchedPath, State},
    http::{self, Request},
    Json, Router,
};
use redis::{aio::MultiplexedConnection, AsyncCommands};
use serde::Deserialize;
use serde_json::{json, Value};
use sqlx::Postgres;
use tower_http::{cors::CorsLayer, trace::TraceLayer};
use tracing::{info, info_span};

use crate::error::ApiResult;

#[derive(Deserialize, Debug)]
struct AxumApiConfig {
    #[serde(rename(deserialize = "api_axum_port"))]
    pub port: u16,
}

#[derive(Clone)]
pub struct AppState {
    pub db: sqlx::Pool<Postgres>,
    pub cache: MultiplexedConnection,
}

#[tracing::instrument(skip_all)]
pub async fn setup(state: AppState) -> Result<()> {
    let config = envy::from_env::<AxumApiConfig>().context("Failed to get env vars")?;

    let app = Router::new()
        .route("/", axum::routing::get(root))
        .layer(
            TraceLayer::new_for_http().make_span_with(|request: &Request<_>| {
                let matched_path = request
                    .extensions()
                    .get::<MatchedPath>()
                    .map(MatchedPath::as_str);

                info_span!(
                    "http_request",
                    method = ?request.method(),
                    matched_path,
                )
            }),
        )
        .layer(CorsLayer::new().allow_headers([http::header::CONTENT_TYPE]))
        .with_state(state.into());

    let listener =
        tokio::net::TcpListener::bind(SocketAddrV4::new(Ipv4Addr::new(0, 0, 0, 0), config.port))
            .await
            .context("Failed to bind to TCP port")?;

    info!("Listening on port {}", config.port);

    axum::serve(listener, app)
        .await
        .context("Axum serve failed")?;

    Ok(())
}

#[tracing::instrument(skip_all)]
async fn root(State(state): State<Arc<AppState>>) -> ApiResult<Json<Value>> {
    let value: (i64,) = sqlx::query_as("SELECT $1")
        .bind(150_i64)
        .fetch_one(&state.db)
        .await?;

    let mut conn = state.cache.clone();

    let _ = &mut conn.set("key", value.0).await?;

    info!("Api logging!");

    Ok(Json(json!(value.0)))
}
