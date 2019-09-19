package services

import (
	"context"

	"github.com/Krystian19/cactus-core/proto"
)

// Release : Get a single Release based on the provided params
func (s *Server) Release(ctx context.Context, request *proto.ReleaseRequest) (*proto.ReleaseResponse, error) {
	return &proto.ReleaseResponse{Release: &proto.Release{Id: 1, Title: "This is a testing title"}}, nil
}

// Releases : Get a single Releases based on the provided params
func (s *Server) Releases(ctx context.Context, request *proto.ReleasesRequest) (*proto.ReleasesResponse, error) {
	return &proto.ReleasesResponse{Releases: []*proto.Release{
		&proto.Release{Id: 1, Title: "This is a testing title"},
	}}, nil
}
