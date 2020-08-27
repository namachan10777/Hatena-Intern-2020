extern crate title_fetcher;

use tonic::transport::Server;
use tonic_health::ServingStatus;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let (mut health_reporter, health_service) = tonic_health::server::health_reporter();
    health_reporter
        .set_service_status("".to_owned(), ServingStatus::Serving)
        .await;

    let addr = "0.0.0.0:50051".parse().unwrap();

    println!("TitleFetcherServer listening on {}", addr);

    Server::builder()
        .add_service(title_fetcher::get_service())
        .add_service(health_service)
        .serve(addr)
        .await?;

    Ok(())
}
