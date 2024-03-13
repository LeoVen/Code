use serde::Deserialize;
use sqlx::postgres::{PgConnectOptions, PgPoolOptions};
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

#[derive(Deserialize, Debug)]
struct TracingConfig {
    #[serde(rename(deserialize = "environment"))]
    pub env: String,
}

#[tokio::main]
async fn main() -> Result<(), sqlx::Error> {
    dotenv::dotenv().ok();

    setup_tracing();
    setup_database().await?;

    info!("Program end");
    Ok(())
}

fn setup_tracing() {
    let env_vars = envy::from_env::<TracingConfig>();
    if let Err(err) = &env_vars {
        let err = err.to_string();
        error!(err, "Error getting env vars for tracing");
    }
    let config = env_vars.unwrap();

    if config.env == "dev" {
        tracing_subscriber::fmt().init();
    } else {
        tracing_subscriber::fmt().json().init();
    }

    info!("Tracing setup finished");
}

#[tracing::instrument]
async fn setup_database() -> Result<(), sqlx::Error> {
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

    let row: (i64,) = sqlx::query_as("SELECT $1")
        .bind(150_i64)
        .fetch_one(&pool)
        .await?;

    let data = row.0;
    info!(data, "Here is the result");

    info!("Database setup finished");

    Ok(())
}
