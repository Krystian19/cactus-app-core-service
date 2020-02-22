package models

import (
	"fmt"
	"os"
	"strings"

	"github.com/jinzhu/gorm"
)

const (
	// DESC : Sort records in descending order
	DESC = "desc"
	// ASC : Sort records in ascending order
	ASC = "asc"
)

// InitDB : Retunrs a db connection object (panics if no connection is stablished)
func InitDB() *gorm.DB {
	// host=myhost port=myport user=gorm dbname=gorm password=mypassword sslmode=disable
	connectionString := fmt.Sprintf(
		"host=%s port=5432 user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
	)

	db, err := gorm.Open("postgres", connectionString)

	if err != nil {
		panic(err)
	}

	// Turn on the logMode to debug db queries
	// db.LogMode(true)

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	db.DB().SetMaxIdleConns(100)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	db.DB().SetMaxOpenConns(1500)

	return db
}

// WhereFieldLikeString : Where field is LIKE "long string with spaces"
func WhereFieldLikeString(query *gorm.DB, field string, value string) *gorm.DB {
	tokens := strings.Fields(strings.TrimSpace(value))

	for _, v := range tokens {
		query = query.Where(fmt.Sprintf("%s ILIKE ?", field), fmt.Sprintf("%%%s%%", v))
	}

	return query
}
