#[macro_use]
extern crate html5ever;
extern crate tonic;

use tonic::{Code, Request, Response, Status};
use pb::title_fetcher_server::{TitleFetcher, TitleFetcherServer};
use pb::{FetchReply, FetchRequest};
use std::io;

pub mod parser;
pub mod pb {
    tonic::include_proto!("title_fetcher");
}

#[derive(Default)]
pub struct TitleFetcherService {}

enum Error {
    HTTP(reqwest::StatusCode),
    Internal(String),
    FailedToSerialize,
}

// TODO test fetch_title
// how to test async function?
async fn fetch_title(url: &str) -> Result<String, Error> {
    let res = reqwest::get(url).await.map_err(|e| {
        e.status()
            .map_or_else(|| Error::Internal(format!("{:?}", e)), Error::HTTP)
    })?;
    let body = res.text().await.map_err(|_| Error::FailedToSerialize)?;
    let title = parser::parse(&mut io::Cursor::new(body));
    Ok(title.unwrap_or_else(String::new))
}

#[tonic::async_trait]
impl TitleFetcher for TitleFetcherService {
    async fn fetch(&self, request: Request<FetchRequest>) -> Result<Response<FetchReply>, Status> {
        println!("Got a request from {:?}", request.remote_addr());
        match fetch_title(&request.into_inner().url).await {
            Ok(title) => Ok(Response::new(pb::FetchReply { title })),
            Err(Error::HTTP(status)) => Err(Status::new(
                Code::InvalidArgument,
                format!("failed to request via HTTP: {:?}", status),
            )),
            Err(Error::Internal(msg)) => Err(Status::new(
                Code::InvalidArgument,
                format!("Invalid argument msg: {}", msg),
            )),
            Err(Error::FailedToSerialize) => Err(Status::new(
                Code::Internal,
                "Internal Error: Failed to serialize text",
            )),
        }
    }
}

pub fn get_service() -> pb::title_fetcher_server::TitleFetcherServer<TitleFetcherService> {
   TitleFetcherServer::new(TitleFetcherService::default())
}
