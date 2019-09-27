package services

import (
	"context"
	"fmt"

	"github.com/Krystian19/cactus-core/models"
	"github.com/Krystian19/cactus-core/proto"
)

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
