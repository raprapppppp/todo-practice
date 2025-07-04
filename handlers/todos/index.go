package todos

import (
	"todo-practice/models/todos"

	"github.com/gofiber/fiber/v2"
)

func AddTodos(c *fiber.Ctx) error {
	payload := new(todos.Todos)

	// parse & instruct the payload
	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Todo added successfully",
		"data":    payload,
	})
}
