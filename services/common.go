package services

import (
	"github.com/jinzhu/gorm"

	// Includes the postgres gorm driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Services : Contains all of the GRPC services methods
type Services struct {
	DB *gorm.DB
}
