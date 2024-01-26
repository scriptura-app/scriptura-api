package db

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	var err error

	DB, err = gorm.Open(postgres.Open(os.Getenv("POSTGRES_URI")), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to the database, retrying in 5 sec...")
		time.Sleep(5 * time.Second)
		Connect()
	}

	return nil
}
