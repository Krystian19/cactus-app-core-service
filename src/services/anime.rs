use anime::anime_service_server::AnimeService;
use tonic::{Request, Response, Status};

#[derive(Debug, Default)]
pub struct Anime;

#[tonic::async_trait]
impl AnimeService for Anime {
  async fn anime(
    &self,
    _request: Request<anime::AnimeRequest>,
  ) -> Result<Response<anime::AnimeResponse>, Status> {
    let res = anime::AnimeResponse {
      anime: Some(anime::Anime {
        id: 1996,
        title: String::from("La cosa esta funcionando"),
        created_at: String::from("Any"),
        updated_at: String::from("Any"),
      }),
    };

    Ok(Response::new(res))
  }
}

pub mod anime {
  tonic::include_proto!("anime");
}