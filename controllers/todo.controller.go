package controllers

import (
	"go-fiber-todos/database"
	"go-fiber-todos/models"

	"github.com/gofiber/fiber/v2"
)

func GetAllTodos(c *fiber.Ctx) error {
	db := database.DBConn

	var todos []models.Todo
	db.Find(&todos)

	return c.Status(fiber.StatusOK).JSON(todos)
}
