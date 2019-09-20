package services

import (
	"context"
	"fmt"
	"log"

	"github.com/Krystian19/cactus-core/models"
	"github.com/Krystian19/cactus-core/proto"
	"github.com/jinzhu/gorm"
)

// Genre : Get a single Genre based on the provided params
func (s *Server) Genre(ctx context.Context, request *proto.GenreRequest) (*proto.GenreResponse, error) {
	db, err := db()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer db.Close()

	var result models.Genre
	query := db

	if request.Id != 0 {
		query = query.Where("id = ?", request.Id)
	}

	if err := query.First(&result).Error; err != nil {

		// If nothing was found
		if gorm.IsRecordNotFoundError(err) {
			return &proto.GenreResponse{Genre: nil}, nil
		}

		fmt.Println(err)
		return nil, err
	}

	return &proto.GenreResponse{Genre: result.Genre}, nil
}

// Genres : Get a list of Genres based on the provided params
func (s *Server) Genres(ctx context.Context, request *proto.GenresRequest) (*proto.GenresResponse, error) {
	db, err := db()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer db.Close()

	var result []models.Genre
	var resultCount uint
	query := db

	if request != nil && request.Query != nil {
		if request.Query.Limit != 0 {
			query = query.Limit(request.Query.Limit)
		}

		if request.Query.Offset != 0 {
			query = query.Offset(request.Query.Offset)
		}
	}

	if err := query.Find(&result).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Remove pagination constraints before counting
	query = query.Limit(nil)
	query = query.Offset(nil)

	if err := query.Model(&models.Genre{}).Count(&resultCount).Error; err != nil {
		resultCount = 0
	}

	finalRes := []*proto.Genre{}

	for i := range result {
		finalRes = append(finalRes, result[i].Genre)
	}

	return &proto.GenresResponse{Genres: finalRes, Count: uint64(resultCount)}, nil
}
