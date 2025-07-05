package todos

import (
	db "todo-practice/database"
	"todo-practice/models/todos"

	"github.com/gofiber/fiber/v2"
)

func AddTodos(c *fiber.Ctx) error {
	payload := new(todos.TodosOrig)

	// parse & instruct the payload
	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	db.Database.Create(&payload)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Todo added successfully",
		"data":    payload,
	})
}

// This is handler request without user id
func GetTask(c *fiber.Ctx) error {
	var tasks []todos.TodosOrig

	db.Database.Find(&tasks)

	return c.JSON(tasks)
}

func DeleteTask(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	var tasks todos.TodosOrig

	result := db.Database.Delete(tasks, id)
	if result.RowsAffected == 0 {
		return c.SendStatus(400)
	}
	return c.SendStatus(200)

}
