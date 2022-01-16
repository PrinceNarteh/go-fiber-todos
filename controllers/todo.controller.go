package controllers

import (
	"go-fiber-todos/database"
	"go-fiber-todos/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllTodos(c *fiber.Ctx) error {
	db := database.DBConn

	var todos []models.Todo
	db.Find(&todos)

	return c.Status(fiber.StatusOK).JSON(todos)
}

func GetTodo(c *fiber.Ctx) error {
	todoId := c.Params("todoId")
	id, err := strconv.Atoi(todoId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Could not parse todoId"})
	}

	db := database.DBConn
	var todo models.Todo

	if err := db.Where("ID = ?", id).First(&todo).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Todo not found"})
	}

	return c.Status(fiber.StatusOK).JSON(todo)

}
