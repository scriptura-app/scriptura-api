package db

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateDBConnection(retryLimit int) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(os.Getenv("POSTGRES_URI")), &gorm.Config{})
	if err != nil && retryLimit > 0 {
		fmt.Println("Failed to connect to the database, retrying in 5 sec...")
		time.Sleep(5 * time.Second)
		db, err = CreateDBConnection(retryLimit - 1)
	}
	return db, err
}
