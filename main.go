package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	PORT := os.Getenv("port")
	if PORT == "" {
		PORT = ":4000"
	}

	app.Get("/health-check", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.Listen(PORT)
}
