package controllers

import (
	"go-fiber-todos/database"
	"go-fiber-todos/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type request struct {
	Name      *string `json:"name"`
	Completed *bool   `json:"completed"`
}

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

	if err := db.First(&todo, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Todo not found"})
	}

	return c.Status(fiber.StatusOK).JSON(todo)
}

func AddTodo(c *fiber.Ctx) error {
	var body request
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse json"})
	}

	todo := models.Todo{
		Name: *body.Name,
	}

	// Insert to DB
	db := database.DBConn
	db.Create(&todo)

	return c.Status(fiber.StatusCreated).JSON(todo)
}

func DeleteTodo(c *fiber.Ctx) error {
	todoId := c.Params("todoId")
	id, err := strconv.Atoi(todoId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "could not parse id"})
	}

	var todo models.Todo
	db := database.DBConn
	if err := db.First(&todo, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Todo not found"})
	}

	// delete todo
	db.Delete(&todo)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "todo deleted successfully."})
}

func UpdateTodo(c *fiber.Ctx) error {
	todoId := c.Params("todoId")
	id, err := strconv.Atoi(todoId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "could not parse id"})
	}

	var body request
	if err = c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "could not parse body"})
	}

	var todo models.Todo
	db := database.DBConn
	if err = db.First(&todo, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Todo not found"})
	}

	if body.Name != nil {
		todo.Name = *body.Name
	}

	if body.Completed != nil {
		todo.Completed = *body.Completed
	}

	db.Save(&todo)
	return c.Status(fiber.StatusOK).JSON(todo)
}
