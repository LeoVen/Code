use serde::Deserialize;
use tracing::info;

use anyhow::{Context, Result};
use redis::{aio::MultiplexedConnection, ConnectionAddr, ConnectionInfo, RedisConnectionInfo};

#[derive(Deserialize, Debug)]
struct RedisConfig {
    #[serde(rename(deserialize = "cache_redis_host"))]
    pub host: String,
    #[serde(rename(deserialize = "cache_redis_port"))]
    pub port: u16,
    #[serde(rename(deserialize = "cache_redis_database"))]
    pub database: i64,
}

pub async fn setup() -> Result<MultiplexedConnection> {
    info!("Setting up redis cache");

    let config = envy::from_env::<RedisConfig>().context("Failed to get env vars")?;

    let client = redis::Client::open(ConnectionInfo {
        addr: ConnectionAddr::Tcp(config.host, config.port),
        redis: RedisConnectionInfo {
            db: config.database,
            username: None,
            password: None,
        },
    })
    .context("Failed to setup redis connection manager")?;

    let conn = client
        .get_multiplexed_tokio_connection()
        .await
        .context("Failed to get redis multiplexed connection")?;

    info!("Redis cache setup finished");

    Ok(conn)
}
