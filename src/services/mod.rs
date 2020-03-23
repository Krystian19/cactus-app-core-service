use async_trait::async_trait;
use std::net::SocketAddr;
use tonic::transport::Server;

use anime::{anime::anime_service_server::AnimeServiceServer, Anime};
mod anime;

pub struct Services;

#[async_trait]
pub trait ServicesMethods {
  async fn build_server(&self, addr: SocketAddr) -> Result<(), Box<dyn std::error::Error>>;
}

#[async_trait]
impl ServicesMethods for Services {
  async fn build_server(&self, addr: SocketAddr) -> Result<(), Box<dyn std::error::Error>> {
    Server::builder()
      .add_service(AnimeServiceServer::new(Anime::default()))
      .serve(addr)
      .await?;

    Ok(())
  }
}
