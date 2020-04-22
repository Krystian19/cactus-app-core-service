use super::{services::proto, Model};
use std::string::String;

pub struct Anime {
  pub id: i64,
  pub title: String,
  pub created_at: std::string::String,
  pub updated_at: std::string::String,
}

impl Model<proto::Anime> for Anime {
  fn to_proto(self) -> proto::Anime {
    proto::Anime {
      id: self.id,
      title: String::from(self.title),
      created_at: String::from(self.created_at),
      updated_at: String::from(self.updated_at),
    }
  }
}
