package services

import (
	"context"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"

	"github.com/Krystian19/cactus-core/models"
	"github.com/Krystian19/cactus-core/proto"
)

// Anime : Get a single Anime based on the provided params
func (s *Server) Anime(ctx context.Context, request *proto.Empty) (*proto.AnimeResponse, error) {
	db, err := db()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer db.Close()

	var anime models.Anime
	if err := db.Where("id = ?", 3).First(&anime).Error; err != nil {

		// If nothing was found
		if gorm.IsRecordNotFoundError(err) {
			return &proto.AnimeResponse{Anime: nil}, nil
		}

		fmt.Println(err)
	}

	return &proto.AnimeResponse{Anime: anime.Anime}, nil
}
