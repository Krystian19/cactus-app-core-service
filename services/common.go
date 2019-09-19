package services

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"

	// Includes the mysql gorm driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Server : Contains all of the services methods
type Server struct{}

// db : Returns a connection to the database
func db() (*gorm.DB, error) {
	return gorm.Open("mysql", fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME")))
}
