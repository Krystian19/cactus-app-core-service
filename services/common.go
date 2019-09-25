package services

import (
	"fmt"
	"os"
	"strings"

	"github.com/jinzhu/gorm"

	// Includes the postgres gorm driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
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
	// host=myhost port=myport user=gorm dbname=gorm password=mypassword
	connectionString := fmt.Sprintf("host=%s port=5432 user=%s dbname=%s password=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_USERNAME"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD"))
	db, err := gorm.Open("postgres", connectionString)

	if err != nil {
		panic(err)
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	db.DB().SetMaxIdleConns(100)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	db.DB().SetMaxOpenConns(1500)

	DB = db
}

// GetDB : Get the current global connection object
func GetDB() *gorm.DB {
	return DB
}

// WhereFieldLikeString : Where field is LIKE "long value string with spaces"
func WhereFieldLikeString(query *gorm.DB, field string, value string) *gorm.DB {
	tokens := strings.Fields(strings.TrimSpace(value))

	for _, v := range tokens {
		query = query.Where(fmt.Sprintf("%s LIKE ?", field), fmt.Sprintf("%%%s%%", v))
	}

	return DB
}
