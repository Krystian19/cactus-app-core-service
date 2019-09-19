package services

import (
	"context"
	"fmt"
	"log"

	"github.com/Krystian19/cactus-core/models"
	"github.com/Krystian19/cactus-core/proto"
	"github.com/jinzhu/gorm"
)

// Release : Get a single Release based on the provided params
func (s *Server) Release(ctx context.Context, request *proto.ReleaseRequest) (*proto.ReleaseResponse, error) {
	db, err := db()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer db.Close()

	var result models.Release
	query := db

	if request.Id != 0 {
		query = query.Where("id = ?", request.Id)
	}

	if err := query.First(&result).Error; err != nil {

		// If nothing was found
		if gorm.IsRecordNotFoundError(err) {
			return &proto.ReleaseResponse{Release: nil}, nil
		}

		fmt.Println(err)
		return nil, err
	}

	return &proto.ReleaseResponse{Release: result.Release}, nil
}

// Releases : Get a single Releases based on the provided params
func (s *Server) Releases(ctx context.Context, request *proto.ReleasesRequest) (*proto.ReleasesResponse, error) {
	db, err := db()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer db.Close()

	var result []models.Release
	query := db

	if request.Query.AnimeId != 0 {
		query = query.Where("anime_id = ?", request.Query.AnimeId)
	}

	if err := query.Find(&result).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	finalRes := []*proto.Release{}

	for i := range result {
    finalRes[i] = result[i].Release
  }

	return &proto.ReleasesResponse{Releases: finalRes}, nil
}
