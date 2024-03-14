use serde::Deserialize;
use sqlx::{
    postgres::{PgConnectOptions, PgPoolOptions},
    Pool, Postgres,
};
use tracing::{error, info};

#[derive(Deserialize, Debug)]
struct DatabasePostgresConfig {
    #[serde(rename(deserialize = "database_postgres_host"))]
    pub host: String,
    #[serde(rename(deserialize = "database_postgres_port"))]
    pub port: u16,
    #[serde(rename(deserialize = "database_postgres_user"))]
    pub user: String,
    #[serde(rename(deserialize = "database_postgres_password"))]
    pub password: String,
}

#[tracing::instrument]
pub async fn setup() -> Result<Pool<Postgres>, sqlx::Error> {
    info!("Setting up database");

    let env_vars = envy::from_env::<DatabasePostgresConfig>();

    if let Err(err) = &env_vars {
        let err = err.to_string();
        error!(err, "Error getting env vars for database");
    }
    let config = env_vars.unwrap();

    let conn_opt = PgConnectOptions::new()
        .host(&config.host)
        .port(config.port)
        .username(&config.user)
        .password(&config.password);

    let pool = PgPoolOptions::new().connect_with(conn_opt).await;

    if let Err(err) = &pool {
        let err = err.to_string();
        error!(err, "Error creating database pool");
    }
    let pool = pool.unwrap();

    info!("Database setup finished");

    Ok(pool)
}
