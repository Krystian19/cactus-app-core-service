use tonic::{Request, Response, Status};
use super::proto::{anime_service_server::AnimeService};
use super::proto;
use crate::models;

#[derive(Debug, Default)]
pub struct Anime;

#[tonic::async_trait]
impl AnimeService for Anime {
  async fn anime(
    &self,
    request: Request<proto::AnimeRequest>, // Accepts requests of type AnimeRequest
  ) -> Result<Response<proto::AnimeReply>, Status> {
    // Return an instance of type HelloReply
    println!("Got a request: {:?}", request);
    let _connection = models::db_connection();

    let reply = proto::AnimeReply {
      anime: Some(
        proto::Anime {
          id: 1996,
          title: String::from("adios"),
          created_at: String::from("adios"),
          updated_at: String::from("adios"),
        }
      )
    };

    Ok(Response::new(reply)) // Send back our formatted greeting
  }
}
