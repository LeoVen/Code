use the_stack::api::AppState;
use tracing::info;

#[tokio::main]
async fn main() -> Result<(), sqlx::Error> {
    dotenv::dotenv().ok();

    the_stack::tracing::setup();
    let pool = the_stack::database::setup().await?;
    let _ = the_stack::api::setup(AppState { pool }).await;

    info!("Program end");
    Ok(())
}
