use the_stack::api::AppState;
use tracing::info;

#[tokio::main]
async fn main() -> anyhow::Result<()> {
    dotenv::dotenv().ok();

    the_stack::tracing::setup();
    let db = the_stack::database::setup().await?;
    let cache = the_stack::cache::setup().await?;
    the_stack::api::setup(AppState { db, cache }).await?;

    info!("Program end");
    Ok(())
}
