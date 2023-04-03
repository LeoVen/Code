use aws_sdk_ec2 as ec2;
use aws_sdk_s3 as s3;
use ec2::types::Subnet;

async fn create_bucket(
    client: &s3::Client,
    bucket_name: impl Into<String>,
) -> Result<(), s3::Error> {
    let bucket = client.create_bucket().bucket(bucket_name).send().await?;

    println!("{:?}", bucket);

    Ok(())
}

async fn list_subnets(client: &ec2::Client) -> Result<Vec<Subnet>, ec2::Error> {
    let subnets = client.describe_subnets().send().await?;

    if let Some(subnets) = subnets.subnets() {
        return Ok(Vec::from(subnets));
    }

    Ok(vec![])
}

#[tokio::main]
async fn main() -> Result<(), ec2::Error> {
    let config = aws_config::from_env().region("us-east-1").load().await;
    let ec2_client = ec2::Client::new(&config);
    let s3_client = s3::Client::new(&config);

    let subnets = list_subnets(&ec2_client).await?;

    let ids: Vec<String> = subnets
        .into_iter()
        .filter(|subnet| subnet.default_for_az().unwrap_or_default())
        .map(|subnet| subnet.subnet_id().unwrap_or_default().to_string())
        .collect();

    let bucket = create_bucket(&s3_client, "hello-world-at-2023").await;

    match bucket {
        Ok(bucket) => println!("Created bucket:\n{:?}", bucket),
        Err(e) => eprintln!("Something went wrong: {}", e),
    }

    println!("{:?}", ids);

    Ok(())
}
