use the_stack::api::AppState;

#[tokio::main]
async fn main() -> anyhow::Result<()> {
    dotenv::dotenv().ok();

    let env = the_stack::tracing::setup();
    let db = the_stack::database::setup(&env).await?;
    let cache = the_stack::cache::setup(&env).await?;
    the_stack::api::setup(&env, AppState { db, cache }).await?;

    tracing::info!("Program end");
    Ok(())
}
