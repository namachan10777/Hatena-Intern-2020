#[macro_use]
extern crate html5ever;
extern crate tokio;
extern crate tonic;

use pb::title_fetcher_server::{TitleFetcher, TitleFetcherServer};
use pb::{FetchReply, FetchRequest};
use std::io;
use std::time;
use regex::Regex;
use tokio::time::delay_for;
use std::collections::HashMap;
use tokio::sync::RwLock;
use tonic::{Code, Request, Response, Status};

const ACCESS_LIMIT: usize = 100;
const DOMAIN_RE_PAT: &'static str = r"(https?://[^\s|^/]+)";

pub mod parser;
pub mod pb {
    tonic::include_proto!("title_fetcher");
}

#[derive(Default)]
pub struct TitleFetcherService {
    cache: RwLock<HashMap<String, (Result<String, Error>, std::time::SystemTime)>>,
    request_count: RwLock<HashMap<String, usize>>,
}

#[derive(PartialEq, Debug, Clone)]
enum Error {
    HTTP(reqwest::StatusCode),
    Internal(String),
    InvalidHttpURL,
    FailedToSerialize,
}

// TODO test fetch_title
// how to test async function?
async fn fetch_title(url: &str) -> Result<String, Error> {
    let client = reqwest::Client::builder()
        .timeout(std::time::Duration::from_millis(1000))
        .build()
        .map_err(|e| Error::Internal(format!("{:?}", e)))?;
    let res = client.get(url).send().await.map_err(|e| {
        e.status()
            .map_or_else(|| Error::Internal(format!("{:?}", e)), Error::HTTP)
    })?;
    if !res.status().is_success() && !res.status().is_redirection() {
        Err(Error::HTTP(res.status()))
    }
    else {
        let body = res.text().await.map_err(|_| Error::FailedToSerialize)?;
        let title = parser::parse(&mut io::Cursor::new(body));
        Ok(title.unwrap_or_else(String::new))
    }
}

impl TitleFetcherService {
    async fn get_title(&self, url: String) -> Result<String, Error> {
        // to free lock early.
        {
            let reader = self.cache.read().await;
            if let Some((cached, timestamp)) = reader.get(&url) {
                if let Ok(dur) = timestamp.elapsed() {
                    if dur > time::Duration::from_secs(10) {
                        return cached.clone()
                    }
                }
            }
            // drop lock
        }
        let domain_re = Regex::new(DOMAIN_RE_PAT).unwrap();
        let domain = domain_re.captures(&url);
        let domain = if let Some(domain) = domain {
            domain.get(0).unwrap().as_str()
        }
        else {
            return Err(Error::InvalidHttpURL)
        };
        let mut wait = 1;
        loop {
            let limit = {
                let reader = self.request_count.read().await;
                reader.get(domain).cloned().unwrap_or(0)
            };
            println!("delay {} for {}", wait, domain);
            if limit < ACCESS_LIMIT {
                let mut writer = self.request_count.write().await;
                writer.insert(domain.to_owned(), limit+1);
                break;
            }
            delay_for(time::Duration::from_secs(wait)).await;
            wait *= 2;
        }
        let res = fetch_title(&url).await;
        let mut writer = self.request_count.write().await;
        let count = writer.get(domain).cloned().unwrap_or(0);
        if count > 0 {
            writer.insert(domain.to_owned(), count-1);
        }
        let mut writer = self.cache.write().await;
        writer.insert(url, (res.clone(), time::SystemTime::now()));
        res
    }
}

#[tonic::async_trait]
impl TitleFetcher for TitleFetcherService {
    async fn fetch(&self, request: Request<FetchRequest>) -> Result<Response<FetchReply>, Status> {
        match self.get_title(request.into_inner().url).await  {
            Ok(title) => Ok(Response::new(FetchReply{ title })),
            Err(Error::HTTP(status)) => Err(Status::new(
                Code::InvalidArgument,
                format!("failed to request via HTTP: {:?}", status),
            )),
            Err(Error::InvalidHttpURL) => Err(Status::new(
                Code::InvalidArgument,
                "Invalid argument msg: invalid HTTP URL".to_owned(),
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

#[cfg(test)]
mod test {
    use super::*;
    #[tokio::test]
    async fn fetch_google() {
        assert_eq!(
            fetch_title("https://google.com").await,
            Ok("Google".to_owned())
        );
    }

    #[tokio::test]
    async fn fetch_404() {
        assert_eq!(
            fetch_title("https://google.com/nowhere").await,
            Err(Error::HTTP(reqwest::StatusCode::from_u16(404).unwrap()))
        );
    }

    #[test]
    fn domain_re() {
        let re = Regex::new(DOMAIN_RE_PAT).unwrap();
        assert_eq!(re.captures("https://google.com/hogehoge/fugafuga").unwrap().get(0).unwrap().as_str(), "https://google.com");
    }
}
