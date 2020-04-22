use super::services;
pub mod anime;

pub trait Model<T> {
  fn to_proto(self) -> T;
}
