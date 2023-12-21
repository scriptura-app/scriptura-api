package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func JsonMiddleware(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)
	return c.Next()
}
