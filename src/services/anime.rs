#[derive(Debug, Default)]
pub struct Anime;

pub mod anime {
  tonic::include_proto!("anime");
}