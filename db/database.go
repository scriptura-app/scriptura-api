package db

import (
	"embed"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

//go:embed db/queries/*
var embedFiles embed.FS

func Connect() error {
	var err error

	db, err = sqlx.Open("postgres", os.Getenv("POSTGRES_URI")+"?sslmode=disable")

	if err != nil {
		return err
	}

	return nil
}

func Select(slice interface{}, queryName string) error {
	q, err := embedFiles.ReadFile(fmt.Sprintf("db/queries/%s.sql", queryName))

	if err != nil {
		panic(err)
	}

	return db.Select(slice, string(q))
}
