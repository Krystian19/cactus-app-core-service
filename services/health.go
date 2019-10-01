package services

import (
	"context"
	"log"

	"google.golang.org/grpc/health/grpc_health_v1"
)

// Check : Responds the healthcheck probe with the current status of the service
func (s *Server) Check(ctx context.Context, request *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	// If Ping was not successful we currently don't have a connection to the database
	if err := s.db.DB().Ping(); err != nil {
		log.Fatal("DataBase Ping failed.............")
		log.Fatal(err)
		return &grpc_health_v1.HealthCheckResponse{
			Status: grpc_health_v1.HealthCheckResponse_NOT_SERVING,
		}, nil
	}

	// All points passed return an all-good response
	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}

// Watch : Performs a watch for the serving status of the requested service. (We are not using it right now)
func (s *Server) Watch(request *grpc_health_v1.HealthCheckRequest, server grpc_health_v1.Health_WatchServer) error {
	return nil
}
