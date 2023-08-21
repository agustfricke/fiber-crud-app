package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/template/html/v2"
	"github.com/agustfricke/fiber-crud-app/handlers"
    "github.com/agustfricke/fiber-crud-app/database"
)

func main() {
    database.ConnectDb()

    engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine, 
        ViewsLayout: "layouts/main", 
	})

	app.Get("/", handlers.ListTasks)
    app.Get("/add", handlers.NewTaskView) 
    app.Post("/add", handlers.CreateTask)
    app.Get("/success", handlers.ConfirmationView)
    app.Get("/task/:id", handlers.ShowTask)
	app.Get("/task/:id/edit", handlers.EditTask)
	app.Patch("/task/:id", handlers.UpdateTask)
    app.Delete("/task/:id", handlers.DeleteTask)
    app.Use(handlers.NotFound)

    app.Static("/", "./public")

	app.Listen(":3000")
}
