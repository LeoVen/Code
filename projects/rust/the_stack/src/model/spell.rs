use redis_macros::{FromRedisValue, ToRedisArgs};
use serde::{Deserialize, Serialize};
use sqlx::FromRow;

use crate::error::{ApiError, ApiResult};

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

    pub async fn get_by_id(db: &sqlx::Pool<sqlx::Postgres>, id: i64) -> ApiResult<Spell> {
        sqlx::query_as("SELECT id, name, damage, created_at, updated_at FROM spell where id = $1")
            .bind(id)
            .fetch_one(db)
            .await
            .map_err(|err| match err {
                sqlx::Error::RowNotFound => ApiError::NotFound(err.to_string()),
                _ => ApiError::Internal(err.into()),
            })
    }

    pub async fn get_all(db: &sqlx::Pool<sqlx::Postgres>) -> ApiResult<Vec<Spell>> {
        let result: Vec<Spell> =
            sqlx::query_as("SELECT id, name, damage, created_at, updated_at FROM spell")
                .fetch_all(db)
                .await?;

        Ok(result)
    }
}
