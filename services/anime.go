package services

import (
	"context"
	"fmt"

	"github.com/Krystian19/cactus-core/models"
	"github.com/Krystian19/cactus-core/proto"
)

// Anime : Get a single Anime based on the provided params
func (s *Server) Anime(ctx context.Context, request *proto.Empty) (*proto.AnimeResponse, error) {
	db, err := db()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	var anime models.Anime
	if err := db.Where("id = ?", 3).First(&anime).Error; err != nil {
		fmt.Println(err)
	}

	return &proto.AnimeResponse{Anime: anime.Anime}, nil
}
