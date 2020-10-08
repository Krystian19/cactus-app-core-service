use super::db;
use async_trait::async_trait;
use std::net::SocketAddr;
use tonic::transport::Server;

use anime::Anime;
use proto::anime_service_server::AnimeServiceServer;
mod anime;

#[derive(Default)]
pub struct Services;

/// Generated code from the proto files.
pub mod proto {
  tonic::include_proto!("proto");
}

#[async_trait]
pub trait ServicesMethods {
  async fn build_server(&self, addr: SocketAddr) -> Result<(), Box<dyn std::error::Error>>;
}

#[async_trait]
impl ServicesMethods for Services {
  async fn build_server(&self, addr: SocketAddr) -> Result<(), Box<dyn std::error::Error>> {
    let pool = db::pool();
    Server::builder()
      .add_service(AnimeServiceServer::new(Anime { pool: pool.clone() }))
      .serve(addr)
      .await?;

    Ok(())
  }
}
