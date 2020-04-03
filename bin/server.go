package main

import (
	"log"
	"net"
	"os"

	"github.com/Krystian19/cactus-core/models"
	"github.com/Krystian19/cactus-core/proto"
	"github.com/Krystian19/cactus-core/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

// Server : Runs a new GRPC server
func Server(port string) error {
	listener, err := net.Listen("tcp", ":"+port)

	if err != nil {
		panic(err)
	}

	EnvVarsCheck()

	srv := grpc.NewServer()
	// Setup the GRPC Server struct (panics if no db connection is stablished)
	registerServices(srv, services.NewServerConfig(models.InitDB()))
	reflection.Register(srv)

	return srv.Serve(listener)
}

func registerServices(srv *grpc.Server, serverConfig *services.Server) {
	proto.RegisterLanguageServiceServer(srv, serverConfig)
	proto.RegisterGenreServiceServer(srv, serverConfig)
	proto.RegisterAnimeServiceServer(srv, serverConfig)
	proto.RegisterReleaseServiceServer(srv, serverConfig)
	proto.RegisterEpisodeServiceServer(srv, serverConfig)
	grpc_health_v1.RegisterHealthServer(srv, serverConfig)
}

// EnvVarsCheck : Checks that important ENV vars are set
func EnvVarsCheck() {
	if os.Getenv("DB_HOST") == "" {
		log.Fatal("DB_HOST env var is not set")
	}

	if os.Getenv("DB_NAME") == "" {
		log.Fatal("DB_NAME env var is not set")
	}

	if os.Getenv("DB_USERNAME") == "" {
		log.Fatal("DB_USERNAME env var is not set")
	}

	if os.Getenv("DB_PASSWORD") == "" {
		log.Fatal("DB_PASSWORD env var is not set")
	}
}
