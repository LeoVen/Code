use serde::Deserialize;

#[derive(Deserialize, Debug)]
pub struct TracingConfig {
    #[serde(rename(deserialize = "environment"))]
    pub env: String,
}

pub fn setup() -> String {
    let env_vars = envy::from_env::<TracingConfig>();
    let config = env_vars.unwrap_or(TracingConfig {
        env: "prod".to_string(),
    });

    if config.env == "dev" {
        tracing_subscriber::fmt()
            .with_max_level(tracing::Level::DEBUG)
            .init();
    } else {
        tracing_subscriber::fmt().json().init();
    }

    tracing::info!("Tracing setup finished");

    config.env
}
