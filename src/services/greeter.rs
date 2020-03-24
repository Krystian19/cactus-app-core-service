use hello_world::{greeter_server::Greeter, HelloReply, HelloRequest, Person};
use tonic::{Request, Response, Status};

#[derive(Debug, Default)]
pub struct MyGreeter {}

#[tonic::async_trait]
impl Greeter for MyGreeter {
  async fn say_hello(
    &self,
    request: Request<HelloRequest>, // Accept request of type HelloRequest
  ) -> Result<Response<HelloReply>, Status> {
    // Return an instance of type HelloReply
    println!("Got a request: {:?}", request);

    let reply = hello_world::HelloReply {
      message: format!("Hello {}!", request.into_inner().name).into(), // We must use .into_inner() as the fields of gRPC requests and responses are private
      person: Some(Person{
        id: 14,
        name: String::from("Jan Guzman"),
      }),
    };

    Ok(Response::new(reply)) // Send back our formatted greeting
  }
}

pub mod hello_world {
  tonic::include_proto!("helloworld"); // The string specified here must match the proto package name
}
