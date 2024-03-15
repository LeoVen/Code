use axum::{
    http::StatusCode,
    response::{IntoResponse, Response},
};
use serde::Serialize;



pub type ApiResult<T> = Result<T, ApiError>;

// Make our own error that wraps `anyhow::Error`.
pub struct ApiError(pub anyhow::Error);

// Tell axum how to convert `ApiError` into a response.
impl IntoResponse for ApiError {
    fn into_response(self) -> Response {
        #[derive(Serialize)]
        struct Body {
            message: String,
            error: String,
        }

        let body = Body {
            message: "Internal Server Error".to_string(),
            error: self.0.to_string(),
        };

        (
            StatusCode::INTERNAL_SERVER_ERROR,
            serde_json::to_string(&body).unwrap_or_default(),
        )
            .into_response()
    }
}

// This enables using `?` on functions that return `Result<_, anyhow::Error>` to turn them into
// `Result<_, ApiError>`. That way you don't need to do that manually.
impl<E> From<E> for ApiError
where
    E: Into<anyhow::Error>,
{
    fn from(err: E) -> Self {
        Self(err.into())
    }
}
