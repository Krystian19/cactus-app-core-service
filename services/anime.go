package services

import (
	"context"

	"github.com/Krystian19/cactus-core/proto"
)

func (s *server) Anime(ctx context.Context, request *proto.Empty) (*proto.Anime, error) {
	return &proto.Anime{Id: 1, Title: "Boku No Hero Academia"}, nil
}
