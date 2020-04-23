use std::net::{IpAddr, Ipv4Addr, SocketAddr};
pub mod db;
pub mod services;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
  const PORT: u16 = 9040;
  println!(
    "GRPC server running @ http://localhost:{}",
    PORT.to_string()
  );
  services::ServicesMethods::build_server(
    &services::Services::default(),
    SocketAddr::new(IpAddr::V4(Ipv4Addr::new(0, 0, 0, 0)), PORT),
  )
  .await?;

  Ok(())
}
