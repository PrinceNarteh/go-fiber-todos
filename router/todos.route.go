package router

import (
	"github.com/gofiber/fiber/v2"

	"go-fiber-todos/controllers"
)

func TodoRoutes(router fiber.Router) {
	router.Get("/", controllers.GetAllTodos)
}
