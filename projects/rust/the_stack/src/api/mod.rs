use std::sync::Arc;

use axum::{
    extract::{MatchedPath, State},
    http::{self, Request},
    Json, Router,
};
use serde_json::{json, Value};
use sqlx::{Pool, Postgres};
use tower_http::{cors::CorsLayer, trace::TraceLayer};
use tracing::{info, info_span};

#[derive(Clone)]
pub struct AppState {
    pub pool: Pool<Postgres>,
}

#[tracing::instrument(skip_all)]
pub async fn setup(state: AppState) -> Result<(), sqlx::Error> {
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

    let listener = tokio::net::TcpListener::bind("0.0.0.0:3000").await.unwrap();
    axum::serve(listener, app).await.unwrap();
    Ok(())
}

#[tracing::instrument(skip_all)]
async fn root(State(state): State<Arc<AppState>>) -> Result<Json<Value>, Json<Value>> {
    info!("Api logging!");

    let result: Result<(i64,), _> = sqlx::query_as("SELECT $1")
        .bind(150_i64)
        .fetch_one(&state.pool)
        .await;

    return match result {
        Err(err) => Result::Err(axum::response::Json(json!(err.to_string()))),
        Ok((val,)) => Result::Ok(axum::response::Json(json!(val))),
    };
}
