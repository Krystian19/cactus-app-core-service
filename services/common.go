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

const (
	// DESC : Sort records in descending order
	DESC = "desc"
	// ASC : Sort records in ascending order
	ASC = "asc"
)

// DB : Global db connection object
var DB *gorm.DB

// InitDB : Setups the global db connection object (panics if no connection is stablished)
func InitDB() {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME")))

	if err != nil {
		panic(err)
	}

	DB = db
}

// GetDB : Get the current global connection object
func GetDB() *gorm.DB {
	return DB
}