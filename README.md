# Cactus Core service

Core service for the Cactus app.

```sh
docker build --no-cache -t cactus-core .
docker run -it -d -v $(pwd):/go/src/app -p 9040:9040 --name cactus-core cactus-core
```

```sh
# Run a GRPC liveness probe
go run github.com/grpc-ecosystem/grpc-health-probe -addr=localhost:9040 -connect-timeout 250ms -rpc-timeout 100ms

# Install and then run
go get github.com/grpc-ecosystem/grpc-health-probe
/go/bin/grpc-health-probe -addr=localhost:9040 -connect-timeout 250ms -rpc-timeout 100ms
```

## License
MIT © [Jan Guzman](https://github.com/Krystian19)
