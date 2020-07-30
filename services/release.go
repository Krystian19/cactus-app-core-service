package services

import (
	"context"
	"fmt"
	"log"

	"github.com/Krystian19/cactus-core/models"
	"github.com/Krystian19/cactus-core/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jinzhu/gorm"
)

// Release : Get a single Release based on the provided params
func (s *Services) Release(ctx context.Context, request *proto.ReleaseRequest) (*proto.ReleaseResponse, error) {
	var result models.Release
	query := s.DB

	if request.Id != 0 {
		query = query.Where("id = ?", request.Id)
	}

	if err := query.First(&result).Error; err != nil {

		// If nothing was found
		if gorm.IsRecordNotFoundError(err) {
			return &proto.ReleaseResponse{Release: nil}, nil
		}

		log.Println(err)
		return nil, err
	}

	return &proto.ReleaseResponse{Release: result.ToProto()}, nil
}

// Releases : Get a list of Releases based on the provided params
func (s *Services) Releases(ctx context.Context, request *proto.ReleasesRequest) (*proto.ReleasesResponse, error) {
	var result []models.Release
	var resultCount int64
	query := s.DB

	if request != nil && request.Query != nil {
		if request.Query.AnimeId != 0 {
			query = query.Where("anime_id = ?", request.Query.AnimeId)
		}

		if request.Query.Title != "" {
			query = models.WhereFieldLikeString(
				query,
				fmt.Sprintf(`"%s".title`, models.Release.TableName(models.Release{})),
				request.Query.Title,
			)
		}

		if len(request.Query.Genres) > 0 {
			// This JOIN method is use of Correlated Subqueries when the foreign key is indexed, (good explanation on the link below)
			// https://www.periscopedata.com/blog/4-ways-to-join-only-the-first-row-in-sql

			// SELECT "Releases".* FROM "Releases" INNER JOIN (
			// 	SELECT * FROM "Releases" AS "Release" WHERE (
			// 		SELECT "release_id" FROM public."ReleaseGenres" WHERE (
			// 			public."ReleaseGenres".genre_id IN (1,4)) AND "Release".id = public."ReleaseGenres".release_id LIMIT 1
			// ) IS NOT NULL) AS "Release" ON public."Releases" .id = "Release".id

			query = query.Joins(
				fmt.Sprintf(`INNER JOIN ( 
					SELECT * FROM "%s" AS "Release" WHERE (
						SELECT "release_id" FROM public."ReleaseGenres" WHERE (
							public."ReleaseGenres".genre_id IN (?)) AND "Release".id = public."ReleaseGenres".release_id LIMIT 1
						) IS NOT NULL) AS "Release" ON public."%s" .id = "Release".id`,
					models.Release.TableName(models.Release{}),
					models.Release.TableName(models.Release{}),
				),
				request.Query.Genres,
			)
		}

		if request.Query.Limit != 0 {
			query = query.Limit(request.Query.Limit)
		}

		if request.Query.Offset != 0 {
			query = query.Offset(request.Query.Offset)
		}
	}

	if err := query.Find(&result).Limit(nil).Offset(nil).Count(&resultCount).Error; err != nil {
		log.Println(err)
		return nil, err
	}

	finalRes := []*proto.Release{}

	for i := range result {
		finalRes = append(finalRes, result[i].ToProto())
	}

	return &proto.ReleasesResponse{Releases: finalRes, Count: resultCount}, nil
}

// AiringReleases : Get a list of AiringReleases based on the provided params
func (s *Services) AiringReleases(ctx context.Context, request *empty.Empty) (*proto.ReleasesListResponse, error) {
	query := s.DB

	var result []models.Release

	query = query.Where("started_airing IS NOT NULL AND stopped_airing IS NULL").Where("release_type_id = ?", 1).Or("release_type_id = ?", 4)
	if err := query.Find(&result).Error; err != nil {
		log.Println(err)
		return nil, err
	}

	finalRes := []*proto.Release{}

	for i := range result {
		finalRes = append(finalRes, result[i].ToProto())
	}

	return &proto.ReleasesListResponse{Releases: finalRes}, nil
}

// RandomRelease : Get a single Random Release
func (s *Services) RandomRelease(ctx context.Context, request *empty.Empty) (*proto.ReleaseResponse, error) {
	var result models.Release
	query := s.DB

	if err := query.Order(gorm.Expr("random()")).First(&result).Error; err != nil {

		// If nothing was found
		if gorm.IsRecordNotFoundError(err) {
			return &proto.ReleaseResponse{Release: nil}, nil
		}

		log.Println(err)
		return nil, err
	}

	return &proto.ReleaseResponse{Release: result.ToProto()}, nil
}

// ReleaseType : Get a single ReleaseType based on the provided params
func (s *Services) ReleaseType(ctx context.Context, request *proto.ReleaseTypeRequest) (*proto.ReleaseTypeResponse, error) {
	var result models.ReleaseType
	query := s.DB

	if request.Id != 0 {
		query = query.Where("id = ?", request.Id)
	}

	if err := query.First(&result).Error; err != nil {

		// If nothing was found
		if gorm.IsRecordNotFoundError(err) {
			return &proto.ReleaseTypeResponse{ReleaseType: nil}, nil
		}

		log.Println(err)
		return nil, err
	}

	return &proto.ReleaseTypeResponse{ReleaseType: result.ToProto()}, nil
}

// ReleaseDescriptions : Get a list of ReleaseDescriptions that belong to the provided release_id
func (s *Services) ReleaseDescriptions(ctx context.Context, request *proto.ReleaseDescriptionsRequest) (*proto.ReleaseDescriptionsResponse, error) {
	var result []models.ReleaseDescription
	query := s.DB

	query = query.Where("release_id = ?", request.ReleaseId)

	if err := query.Find(&result).Error; err != nil {
		log.Println(err)
		return nil, err
	}

	finalRes := []*proto.ReleaseDescription{}

	for i := range result {
		finalRes = append(finalRes, result[i].ToProto())
	}

	return &proto.ReleaseDescriptionsResponse{ReleaseDescriptions: finalRes}, nil
}
