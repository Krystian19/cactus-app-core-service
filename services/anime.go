package services

import (
	"context"

	"github.com/Krystian19/cactus-core/proto"
)

// Anime : Get a single Anime based on the provided params
func (s *Server) Anime(ctx context.Context, request *proto.Empty) (*proto.AnimeResponse, error) {
	return &proto.AnimeResponse{Anime: &proto.Anime{Id: 1, Title: "Boku No Hero Academia"}}, nil
}
