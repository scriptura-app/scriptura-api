package handler

import (
	"encoding/json"
	"log"
	"scriptura/scriptura-api/db"
	"scriptura/scriptura-api/models"

	"github.com/gofiber/fiber/v2"
)

func GetVerse(c *fiber.Ctx) error {
	db := db.DB
	var verse models.Verse

	db.First(&verse)

	jsonData, err := json.Marshal(verse)

	if err != nil {
		log.Fatal(err)
	}

	return c.SendString(string(jsonData))
}
