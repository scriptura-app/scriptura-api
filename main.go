package main

import (
	"log"
	"scriptura/scriptura-api/db"
	"scriptura/scriptura-api/router"

	"github.com/gofiber/fiber/v2"
)

func main() {

	db.Connect()
	app := fiber.New()

	router.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
