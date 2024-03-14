use serde::Deserialize;
use tracing::info;

#[derive(Deserialize, Debug)]
pub struct TracingConfig {
    #[serde(rename(deserialize = "environment"))]
    pub env: String,
}

pub fn setup() {
    let env_vars = envy::from_env::<TracingConfig>();
    let config = env_vars.unwrap_or(TracingConfig {
        env: "prod".to_string(),
    });

    if config.env == "dev" {
        tracing_subscriber::fmt().init();
    } else {
        tracing_subscriber::fmt().json().init();
    }

    info!("Tracing setup finished");
}
