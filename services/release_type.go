package services

import (
	"context"
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/Krystian19/cactus-core/models"
	"github.com/Krystian19/cactus-core/proto"
)

// ReleaseType : Get a single ReleaseType based on the provided params
func (s *Server) ReleaseType(ctx context.Context, request *proto.ReleaseTypeRequest) (*proto.ReleaseTypeResponse, error) {
	db := GetDB()

	var result models.ReleaseType
	query := db

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