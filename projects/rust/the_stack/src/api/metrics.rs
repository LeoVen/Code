use axum::Router;
use prometheus::{Encoder, TextEncoder};

use crate::error::ApiResult;

pub fn router() -> Router {
    Router::new().route("/metrics", axum::routing::get(metrics))
}

async fn metrics() -> ApiResult<String> {
    let registry = prometheus::default_registry();

    let mut buffer = vec![];
    let encoder = TextEncoder::new();
    let metric_families = registry.gather();

    encoder.encode(&metric_families, &mut buffer)?;

    Ok(String::from_utf8(buffer)?)
}
