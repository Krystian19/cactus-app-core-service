package services

import (
	"context"
	"fmt"

	"github.com/Krystian19/cactus-core/models"
	"github.com/Krystian19/cactus-core/proto"
	"github.com/jinzhu/gorm"
)

// Release : Get a single Release based on the provided params
func (s *Server) Release(ctx context.Context, request *proto.ReleaseRequest) (*proto.ReleaseResponse, error) {
	var result models.Release
	query := s.db

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

// Releases : Get a list of Releases based on the provided params
func (s *Server) Releases(ctx context.Context, request *proto.ReleasesRequest) (*proto.ReleasesResponse, error) {
	var result []models.Release
	var resultCount uint
	query := s.db

	query = query.Table("Releases")

	if request != nil && request.Query != nil {
		if request.Query.AnimeId != 0 {
			query = query.Where("anime_id = ?", request.Query.AnimeId)
		}

		if request.Query.Title != "" {
			query = models.WhereFieldLikeString(query, "\"Releases\".title", request.Query.Title)
		}

		if len(request.Query.Genres) > 0 {
			// This JOIN method is use of Correlated Subqueries when the foreign key is indexed, (good explanation on the link below)
			// https://www.periscopedata.com/blog/4-ways-to-join-only-the-first-row-in-sql

			// SELECT "Releases".* FROM "Releases" INNER JOIN (
			// 	SELECT * FROM "Releases" AS "Release" WHERE (
			// 		SELECT "release_id" FROM public."ReleaseGenres" WHERE (public."ReleaseGenres".genre_id IN (1,4)) AND "Release".id = public."ReleaseGenres".release_id LIMIT 1
			// ) IS NOT NULL) AS "Release" ON public."Releases" .id = "Release".id

			query = query.Joins("INNER JOIN ( SELECT * FROM \"Releases\" AS \"Release\" WHERE ( SELECT \"release_id\" FROM public.\"ReleaseGenres\" WHERE (public.\"ReleaseGenres\".genre_id IN (?)) AND \"Release\".id = public.\"ReleaseGenres\".release_id LIMIT 1 ) IS NOT NULL) AS \"Release\" ON public.\"Releases\" .id = \"Release\".id", request.Query.Genres)
		}

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

	if err := query.Model(&models.Release{}).Count(&resultCount).Error; err != nil {
		resultCount = 0
	}

	finalRes := []*proto.Release{}

	for i := range result {
		finalRes = append(finalRes, result[i].Release)
	}

	return &proto.ReleasesResponse{Releases: finalRes, Count: uint64(resultCount)}, nil
}

// AiringReleases : Get a list of AiringReleases based on the provided params
func (s *Server) AiringReleases(ctx context.Context, request *proto.Empty) (*proto.ReleasesListResponse, error) {
	db := s.db

	var result []models.Release

	if err := db.Where("started_airing IS NOT NULL AND stopped_airing IS NULL").Where("release_type_id = ?", 1).Or("release_type_id = ?", 4).Find(&result).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	finalRes := []*proto.Release{}

	for i := range result {
		finalRes = append(finalRes, result[i].Release)
	}

	return &proto.ReleasesListResponse{Releases: finalRes}, nil
}

// RandomRelease : Get a single Random Release
func (s *Server) RandomRelease(ctx context.Context, request *proto.Empty) (*proto.ReleaseResponse, error) {
	var result models.Release
	query := s.db

	if err := query.Order(gorm.Expr("random()")).First(&result).Error; err != nil {

		// If nothing was found
		if gorm.IsRecordNotFoundError(err) {
			return &proto.ReleaseResponse{Release: nil}, nil
		}

		fmt.Println(err)
		return nil, err
	}

	return &proto.ReleaseResponse{Release: result.Release}, nil
}

// ReleaseType : Get a single ReleaseType based on the provided params
func (s *Server) ReleaseType(ctx context.Context, request *proto.ReleaseTypeRequest) (*proto.ReleaseTypeResponse, error) {
	var result models.ReleaseType
	query := s.db

	if request.Id != 0 {
		query = query.Where("id = ?", request.Id)
	}

	if err := query.First(&result).Error; err != nil {

		// If nothing was found
		if gorm.IsRecordNotFoundError(err) {
			return &proto.ReleaseTypeResponse{ReleaseType: nil}, nil
		}

		fmt.Println(err)
		return nil, err
	}

	return &proto.ReleaseTypeResponse{ReleaseType: result.ReleaseType}, nil
}

// ReleaseDescriptions : Get a list of ReleaseDescriptions that belong to the provided release_id
func (s *Server) ReleaseDescriptions(ctx context.Context, request *proto.ReleaseDescriptionsRequest) (*proto.ReleaseDescriptionsResponse, error) {
	var result []models.ReleaseDescription
	query := s.db

	query = query.Where("release_id = ?", request.ReleaseId)

	if err := query.Find(&result).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	finalRes := []*proto.ReleaseDescription{}

	for i := range result {
		finalRes = append(finalRes, result[i].ReleaseDescription)
	}

	return &proto.ReleaseDescriptionsResponse{ReleaseDescriptions: finalRes}, nil
}
