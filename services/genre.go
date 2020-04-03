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
func (s *Services) Genre(ctx context.Context, request *proto.GenreRequest) (*proto.GenreResponse, error) {
	var result models.Genre
	query := s.DB

	if request.Id != 0 {
		query = query.Where("id = ?", request.Id)
	}

	if err := query.First(&result).Error; err != nil {

		// If nothing was found
		if gorm.IsRecordNotFoundError(err) {
			return &proto.GenreResponse{Genre: nil}, nil
		}

		log.Println(err)
		return nil, err
	}

	return &proto.GenreResponse{Genre: result.Genre}, nil
}

// Genres : Get a list of Genres based on the provided params
func (s *Services) Genres(ctx context.Context, request *proto.GenresRequest) (*proto.GenresResponse, error) {
	var result []models.Genre
	var resultCount int64
	query := s.DB

	if request != nil && request.Query != nil {
		if request.Query.Title != "" {
			query = models.WhereFieldLikeString(
				query,
				fmt.Sprintf(`"%s".title`, models.Genre.TableName(models.Genre{})),
				request.Query.Title,
			)
		}

		if request.Query.Limit != 0 {
			query = query.Limit(request.Query.Limit)
		}

		if request.Query.Offset != 0 {
			query = query.Offset(request.Query.Offset)
		}
	}

	if err := query.Find(&result).Error; err != nil {
		log.Println(err)
		return nil, err
	}

	// Remove pagination constraints before counting
	query = query.Limit(nil)
	query = query.Offset(nil)

	if err := query.Model(&models.Genre{}).Count(&resultCount).Error; err != nil {
		log.Println(err)
		return nil, err
	}

	finalRes := []*proto.Genre{}

	for i := range result {
		finalRes = append(finalRes, result[i].Genre)
	}

	return &proto.GenresResponse{Genres: finalRes, Count: resultCount}, nil
}

// ReleaseGenres : Get a list of Genres associated with the specified ReleaseId
func (s *Services) ReleaseGenres(ctx context.Context, request *proto.ReleaseGenresRequest) (*proto.GenresListResponse, error) {
	var result []models.Genre
	query := s.DB

	// INNER JOIN public."ReleaseGenres" ON "Genres" .id = public."ReleaseGenres".genre_id WHERE (release_id = 1)
	query = query.Joins(
		fmt.Sprintf(
			`INNER JOIN public."ReleaseGenres" ON "%s" .id = public."ReleaseGenres".genre_id`,
			models.Genre.TableName(models.Genre{}),
		),
	)

	query = query.Where("release_id = ?", request.ReleaseId)

	if err := query.Find(&result).Error; err != nil {
		log.Println(err)
		return nil, err
	}

	finalRes := []*proto.Genre{}

	for i := range result {
		finalRes = append(finalRes, result[i].Genre)
	}

	return &proto.GenresListResponse{Genres: finalRes}, nil
}
