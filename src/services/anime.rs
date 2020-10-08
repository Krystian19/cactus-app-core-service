use super::db::{models::anime, models::Model};
use super::proto;
use super::proto::anime_service_server::AnimeService;
use super::db::PgPool;
use tonic::{Request, Response, Status};
extern crate diesel;

pub struct Anime {
  pub pool: PgPool,
}

#[tonic::async_trait]
impl AnimeService for Anime {
  async fn anime(
    &self,
    request: Request<proto::AnimeRequest>, // Accepts requests of type AnimeRequest
  ) -> Result<Response<proto::AnimeReply>, Status> {
    println!("Got a request: {:?}", request);
    let _pool = &self.pool;

    // https://github.com/diesel-rs/diesel/issues/2232#issuecomment-580744580
    // pool.get().
    // diesel::


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
