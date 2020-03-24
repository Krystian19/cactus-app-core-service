# Rust GRPC server demo

Hello world rust GRPC server

## Before running you need this
```sh
cargo install --version 1.12.0 watchexec
```

## Run the project
```sh
make
```

```sh
docker build -f Dockerfile.dev -t krystian19/rust_server .
docker run -ti --name=rust_server -d -v $(pwd):/app -p 9040:9040 krystian19/rust_server
docker exec -ti rust_server /bin/bash
```