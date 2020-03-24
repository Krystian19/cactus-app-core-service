use tonic::{Request, Response, Status};

use proto::{anime_service_server::AnimeService, AnimeReply, AnimeRequest};

#[derive(Debug, Default)]
pub struct Anime {}

#[tonic::async_trait]
impl AnimeService for Anime {
  async fn anime(
    &self,
    request: Request<AnimeRequest>, // Accept request of type AnimeRequest
  ) -> Result<Response<AnimeReply>, Status> {
    // Return an instance of type HelloReply
    println!("Got a request: {:?}", request);

    let reply = AnimeReply {
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

pub mod proto {
  tonic::include_proto!("proto");
}
