package main

import (
	"github.com/gofiber/fiber/v2"
	"go-todo-fiber/handlers"
)

func main() {
	app := fiber.New()

	handlers.LoadTasks()

	app.Static("/", "./static")

	app.Get("/api/tasks", handlers.GetTasks)
	app.Post("/api/tasks", handlers.AddTask)
	app.Post("/api/tasks/:id/done", handlers.MarkTaskDone)

	app.Listen(":3000")
}
