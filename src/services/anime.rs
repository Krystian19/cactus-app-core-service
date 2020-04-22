use super::proto;
use super::proto::anime_service_server::AnimeService;
use crate::db::{models::anime, models::Model, PgPool};
use tonic::{Request, Response, Status};

pub struct Anime {
  pub db: PgPool,
}

#[tonic::async_trait]
impl AnimeService for Anime {
  async fn anime(
    &self,
    request: Request<proto::AnimeRequest>, // Accepts requests of type AnimeRequest
  ) -> Result<Response<proto::AnimeReply>, Status> {
    // Return an instance of type HelloReply
    println!("Got a request: {:?}", request);
    // println!("Tu nombre es {:?}", self.nombre);
    // let _connection = db::db_connection();
    // let _connection = self::db;
    let _connection = &self.db;

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
