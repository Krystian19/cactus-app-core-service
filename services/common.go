package services

import (
	"github.com/jinzhu/gorm"

	// Includes the postgres gorm driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Server : Contains all of the services methods
type Server struct {
	db *gorm.DB
}

// NewServerConfig : Returns a pointer to a new Server struct
func NewServerConfig(db *gorm.DB) *Server {
	return &Server{db: db}
}
