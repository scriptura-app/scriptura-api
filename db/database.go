package db

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() error {
	var err error

	DB, err = sql.Open("postgres", os.Getenv("POSTGRES_URI"))
	if err != nil {
		return err
	}
	if err = DB.Ping(); err != nil {
		return err
	}
	return nil
}
