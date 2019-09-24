package main

import (
	"github.com/Krystian19/cactus-core/proto"
	"github.com/Krystian19/cactus-core/services"
	"google.golang.org/grpc"
)

func registerServices(srv *grpc.Server) {
	proto.RegisterGenreServiceServer(srv, &services.Server{})
	proto.RegisterAnimeServiceServer(srv, &services.Server{})
	proto.RegisterReleaseServiceServer(srv, &services.Server{})
	proto.RegisterReleaseTypeServiceServer(srv, &services.Server{})
	proto.RegisterEpisodeServiceServer(srv, &services.Server{})
}
