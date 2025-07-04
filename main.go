package main

import (
	db "todo-practice/database"
	"todo-practice/handlers/todos"
	"todo-practice/handlers/users"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db.ConnectionDB()

	app := fiber.New()

	app.Post("/todos", todos.AddTodos)
	app.Post("/users", users.AddUser)

	app.Listen(":3000")

}
