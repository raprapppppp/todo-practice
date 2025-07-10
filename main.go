package main

import (
	db "todo-practice/database"
	"todo-practice/handlers/todos"
	"todo-practice/handlers/users"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	db.ConnectionDB()

	app := fiber.New()
	app.Use(cors.New())

	app.Post("/users", users.AddUser)

	app.Post("/todos", todos.AddTodos)
	app.Get("/todos", todos.GetTask)
	app.Delete("/todos/:id", todos.DeleteTask)
	app.Put("/todos", todos.UpdateTask)

	app.Listen(":3001")

}
