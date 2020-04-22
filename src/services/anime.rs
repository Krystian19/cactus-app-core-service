use super::db::{models::anime, models::Model};
use super::proto;
use super::proto::anime_service_server::AnimeService;
use super::Service;
use tonic::{Request, Response, Status};

pub struct Anime {
  pub service: Service,
}

#[tonic::async_trait]
impl AnimeService for Anime {
  async fn anime(
    &self,
    request: Request<proto::AnimeRequest>, // Accepts requests of type AnimeRequest
  ) -> Result<Response<proto::AnimeReply>, Status> {
    println!("Got a request: {:?}", request);
    let _connection = &self.service.db;

    let result = anime::Anime {
      id: 14,
      title: String::from("Adios"),
      created_at: String::from("Adios"),
      updated_at: String::from("Adios"),
    };

    let reply = proto::AnimeReply {
      anime: Some(result.to_proto()),
    };

    Ok(Response::new(reply))
  }
}
