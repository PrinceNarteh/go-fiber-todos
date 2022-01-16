package router

import (
	"github.com/gofiber/fiber/v2"

	"go-fiber-todos/controllers"
)

func TodoRoutes(router fiber.Router) {
	router.Get("/", controllers.GetAllTodos)
	router.Post("/", controllers.AddTodo)
	router.Get("/:todoId", controllers.GetTodo)
	router.Patch("/:todoId", controllers.UpdateTodo)
	router.Delete("/:todoId", controllers.DeleteTodo)
}
