use anyhow::{Context, Result};
use serde::Deserialize;
use sqlx::{
    postgres::{PgConnectOptions, PgPoolOptions},
    Pool, Postgres,
};
use tracing::info;

#[derive(Deserialize, Debug)]
struct DatabasePostgresConfig {
    #[serde(rename(deserialize = "database_postgres_host"))]
    pub host: String,
    #[serde(rename(deserialize = "database_postgres_port"))]
    pub port: u16,
    #[serde(rename(deserialize = "database_postgres_username"))]
    pub username: String,
    #[serde(rename(deserialize = "database_postgres_password"))]
    pub password: String,
    #[serde(rename(deserialize = "database_postgres_database"))]
    pub database: String,
}

#[tracing::instrument]
pub async fn setup() -> Result<Pool<Postgres>> {
    info!("Setting up database");

    let config = envy::from_env::<DatabasePostgresConfig>().context("Failed to get env vars")?;

    let conn_opt = PgConnectOptions::new()
        .host(&config.host)
        .port(config.port)
        .username(&config.username)
        .password(&config.password)
        .database(&config.database);

    let pool = PgPoolOptions::new()
        .connect_with(conn_opt)
        .await
        .context("Failed to connect to PostgreSQL")?;

    info!("Database setup finished");

    Ok(pool)
}
