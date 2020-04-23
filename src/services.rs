use super::db;
use super::db::PgPool;
use async_trait::async_trait;
use std::net::SocketAddr;
use tonic::transport::Server;

use anime::Anime;
use proto::anime_service_server::AnimeServiceServer;
mod anime;

#[derive(Default)]
pub struct Services;

pub struct Service {
  pub db: PgPool,
}

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
    let service = Service {
      db: db::db_connection(),
    };
    Server::builder()
      .add_service(AnimeServiceServer::new(Anime { service }))
      .serve(addr)
      .await?;

    Ok(())
  }
}
