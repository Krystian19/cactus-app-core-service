package services

import (
	"context"
	"log"

	"github.com/jinzhu/gorm"

	"github.com/Krystian19/cactus-core/models"
	"github.com/Krystian19/cactus-core/proto"
)

// Language : Get a single Language based on the provided params
func (s *Services) Language(ctx context.Context, request *proto.LanguageRequest) (*proto.LanguageResponse, error) {
	var result models.Language
	query := s.DB

	if request.Id != 0 {
		query = query.Where("id = ?", request.Id)
	}

	if err := query.First(&result).Error; err != nil {

		// If nothing was found
		if gorm.IsRecordNotFoundError(err) {
			return &proto.LanguageResponse{Language: nil}, nil
		}

		log.Println(err)
		return nil, err
	}

	return &proto.LanguageResponse{Language: result.ToProto()}, nil
}
