# Cactus Core service
[![Build Status](https://travis-ci.org/Krystian19/cactus-app-core-service.svg?branch=master)](https://travis-ci.org/Krystian19/cactus-app-core-service) [![Codacy Badge](https://app.codacy.com/project/badge/Grade/108d242cd8da436c8634c1760859fa1d)](https://www.codacy.com/gh/Krystian19/cactus-app-core-service/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=Krystian19/cactus-app-core-service&amp;utm_campaign=Badge_Grade)

Core service for the Cactus app.

```sh
# Run a GRPC liveness probe
go run github.com/grpc-ecosystem/grpc-health-probe -addr=localhost:9040 -connect-timeout 250ms -rpc-timeout 100ms

# Install and then run
go get github.com/grpc-ecosystem/grpc-health-probe
/go/bin/grpc-health-probe -addr=localhost:9040 -connect-timeout 250ms -rpc-timeout 100ms
```

## License
MIT © [Jan Guzman](https://github.com/Krystian19)
