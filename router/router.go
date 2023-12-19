package router

import (
	"scriptura/scriptura-api/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app fiber.Router) {
	api := app.Group("/api", logger.New())

	//temp inline json middleware
	api.Use(func(c *fiber.Ctx) error {
		c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)
		return c.Next()
	})

	verse := api.Group("/bible/:book?/:chapter?/:verse?")
	verse.Get("/", handler.GetVerse)
}
