package main

import (
	"log"
	"scriptura/scriptura-api/db"
	"scriptura/scriptura-api/gql"
	"scriptura/scriptura-api/router"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	err := db.Connect()

	if err != nil {
		panic(err)
	}

	err = gql.CreateSchema()

	if err != nil {
		panic(err)
	}

	app := fiber.New()

	router.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
