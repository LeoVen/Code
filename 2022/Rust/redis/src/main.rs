use std::error::Error;

use redis::Client;

#[tokio::main]
async fn main() -> Result<(), Box<dyn Error>> {

    let redis_client = Client::open("redis://127.0.0.1/")?;
    let mut con = redis_client.get_tokio_connection().await?;

    println!("End");

    Ok(())
}
