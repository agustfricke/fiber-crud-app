package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/agustfricke/fiber-crud-app/handlers"

)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.ListTasks)
    app.Get("/add", handlers.NewTaskView) 
    app.Post("/add", handlers.CreateTask)
    app.Get("/success", handlers.ConfirmationView)
    app.Get("/task/:id", handlers.ShowTask)
	app.Get("/task/:id/edit", handlers.EditTask)
	app.Patch("/task/:id", handlers.UpdateTask)
    app.Delete("/task/:id", handlers.DeleteTask)
    app.Use(handlers.NotFound)
}


