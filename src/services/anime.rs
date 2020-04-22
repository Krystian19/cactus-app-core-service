use tonic::{Request, Response, Status};
use super::proto::{anime_service_server::AnimeService};
use super::proto;
use crate::db;

pub struct Anime {
  pub db: db::PgPool,
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
