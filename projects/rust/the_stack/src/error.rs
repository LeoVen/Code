use axum::{
    http::StatusCode,
    response::{IntoResponse, Response},
};
use serde::Serialize;

pub type ApiResult<T> = Result<T, ApiError>;

// Make our own error that wraps `anyhow::Error`.
pub enum ApiError {
    Internal(anyhow::Error),
    NotFound(String),
}

#[derive(Serialize)]
struct ResponseBody {
    message: String,
}

impl ResponseBody {
    pub fn from(message: &str) -> String {
        serde_json::to_string(&Self {
            message: message.to_string(),
        })
        .unwrap_or_default()
    }
}

// Tell axum how to convert `ApiError` into a response.
impl IntoResponse for ApiError {
    fn into_response(self) -> Response {
        match self {
            ApiError::Internal(error) => {
                let error = error.to_string();
                tracing::error!(error, "Internal Server Error");

                (
                    StatusCode::INTERNAL_SERVER_ERROR,
                    ResponseBody::from("Internal Server Error"),
                )
                    .into_response()
            }
            ApiError::NotFound(message) => {
                tracing::error!(error = message, "Not Found");

                (StatusCode::NOT_FOUND, ResponseBody::from("Not Found")).into_response()
            }
        }
    }
}

// This enables using `?` on functions that return `Result<_, anyhow::Error>` to turn them into
// `Result<_, ApiError>`. That way you don't need to do that manually.
impl<E> From<E> for ApiError
where
    E: Into<anyhow::Error>,
{
    fn from(err: E) -> Self {
        Self::Internal(err.into())
    }
}
