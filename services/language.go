package services

import (
	"context"
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/Krystian19/cactus-core/models"
	"github.com/Krystian19/cactus-core/proto"
)

// Language : Get a single Language based on the provided params
func (s *Server) Language(ctx context.Context, request *proto.LanguageRequest) (*proto.LanguageResponse, error) {
	var result models.Language
	query := s.db

	if request.Id != 0 {
		query = query.Where("id = ?", request.Id)
	}

	if err := query.First(&result).Error; err != nil {

		// If nothing was found
		if gorm.IsRecordNotFoundError(err) {
			return &proto.LanguageResponse{Language: nil}, nil
		}

		fmt.Println(err)
		return nil, err
	}

	return &proto.LanguageResponse{Language: result.Language}, nil
}
