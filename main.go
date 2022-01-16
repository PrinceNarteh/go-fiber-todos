package main

import (
	"os"

	"go-fiber-todos/database"
	"go-fiber-todos/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// connecting to database
	database.Connect()

	// instantiating fiber app
	app := fiber.New()
	PORT := os.Getenv("port")
	if PORT == "" {
		PORT = ":4000"
	}

	app.Get("/health-check", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	routes := app.Group("/api/todos")
	router.TodoRoutes(routes)

	// listen/Serve the new Fiber app
	err := app.Listen(PORT)

	// handle panic errors => panic built-in function that stops the execution of a function and immediately normal execution of that function with an error
	if err != nil {
		panic(err)
	}
}
