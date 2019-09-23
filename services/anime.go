package services

import (
	"context"
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/Krystian19/cactus-core/models"
	"github.com/Krystian19/cactus-core/proto"
)

// Anime : Get a single Anime based on the provided params
func (s *Server) Anime(ctx context.Context, request *proto.AnimeRequest) (*proto.AnimeResponse, error) {
	db := GetDB()

	var result models.Anime
	query := db

	if request.Id != 0 {
		query = query.Where("id = ?", request.Id)
	}

	if err := query.First(&result).Error; err != nil {

		// If nothing was found
		if gorm.IsRecordNotFoundError(err) {
			return &proto.AnimeResponse{Anime: nil}, nil
		}

		fmt.Println(err)
		return nil, err
	}

	return &proto.AnimeResponse{Anime: result.Anime}, nil
}
