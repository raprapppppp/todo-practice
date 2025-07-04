package users

import (
	db "todo-practice/database"
	"todo-practice/models/users"

	"github.com/gofiber/fiber/v2"
)

func AddUser(c *fiber.Ctx) error {

	payload := new(users.Users)

	// parse & instruct the payload
	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	// Check if username exist if exist throw err if not proceed to create
	existingUser := db.Database.Where("username = ? ", payload.Username).First(&payload).Error

	if existingUser == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"username": "Username already exist"})
	}

	// Create user & handler create user error
	newUser := db.Database.Create(&payload).Error

	if newUser != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to create user",
			"message": newUser.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Users added successfully",
		"data":    payload,
	})

}
