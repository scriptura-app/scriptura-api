package handler

import (
	"encoding/json"
	"log"
	"os"
	"scriptura/scriptura-api/db"

	"github.com/gofiber/fiber/v2"
)

type Result struct {
	VerseNum int    `json:"verse_num"`
	Text     string `json:"text"`
}

func GetVerse(c *fiber.Ctx) error {
	// id := c.Params("id")
	db := db.DB
	query, err := os.ReadFile("db/queries/select_verse.sql")
	if err != nil {
		log.Fatal(err)
	}
	var result []Result

	db.Raw(string(query)).Scan(&result)

	jsonData, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

	return c.SendString(string(jsonData))
}
