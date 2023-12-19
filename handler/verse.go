package handler

import (
	"fmt"
	"log"
	"os"
	"scriptura/scriptura-api/db"

	"github.com/gofiber/fiber/v2"
)

func GetVerse(c *fiber.Ctx) error {
	// id := c.Params("id")
	db := db.DB
	query, err := os.ReadFile("db/queries/select_verse.sql")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query(string(query))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rows)
	return c.SendString("hey")
}