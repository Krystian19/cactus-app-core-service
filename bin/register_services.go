package main

import (
	"github.com/Krystian19/cactus-core/proto"
	"github.com/Krystian19/cactus-core/services"
	"google.golang.org/grpc"
)

func registerServices(srv *grpc.Server) {
	server := services.NewServer()

	proto.RegisterAnimeServiceServer(srv, server)
}
