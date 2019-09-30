package main

import (
	"github.com/Krystian19/cactus-core/proto"
	"github.com/Krystian19/cactus-core/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func registerServices(srv *grpc.Server, serverConfig *services.Server) {
	proto.RegisterLanguageServiceServer(srv, serverConfig)
	proto.RegisterGenreServiceServer(srv, serverConfig)
	proto.RegisterAnimeServiceServer(srv, serverConfig)
	proto.RegisterReleaseServiceServer(srv, serverConfig)
	proto.RegisterEpisodeServiceServer(srv, serverConfig)
	grpc_health_v1.RegisterHealthServer(srv, serverConfig)
}
