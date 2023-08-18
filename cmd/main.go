package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/template/html/v2"
    "github.com/agustfricke/fiber-crud-app/database"
)

func main() {
    database.ConnectDb()

    engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine, 
        ViewsLayout: "layouts/main", 
	})

    setupRoutes(app)

    app.Static("/", "./public")

	app.Listen(":3000")
}
