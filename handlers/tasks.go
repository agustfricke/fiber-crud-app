package handlers

import (
    "github.com/gofiber/fiber/v2"
    "github.com/agustfricke/fiber-crud-app/models"
    "github.com/agustfricke/fiber-crud-app/database"
)

func ListTasks(c *fiber.Ctx) error {
	tasks := []models.Task{}
	database.DB.Db.Find(&tasks)
	return c.Render("index", fiber.Map{
		"Body": "Test body",
		"Completed": "Completed test!",
        "Tasks":    tasks,
	})
}

func NewTaskView(c *fiber.Ctx) error {
	return c.Render("new", fiber.Map{
		"Title":    "New Task",
		"Subtitle": "Add a new task!",
	})
}

func CreateTask(c *fiber.Ctx) error {
	task := new(models.Task)
	if err := c.BodyParser(task); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	database.DB.Db.Create(&task)
    return ConfirmationView(c) 
}


func ConfirmationView(c *fiber.Ctx) error {
	return c.Render("confirmation", fiber.Map{
		"Title":    "Fact added successfully",
		"Subtitle": "Add more wonderful tasks to the list!",
	})
}

func ShowTask(c *fiber.Ctx) error {
	task := models.Task{}
	id := c.Params("id")

    result := database.DB.Db.Where("id = ?", id).First(&task)
    if result.Error != nil {
		return NotFound(c)
	}

	return c.Render("show", fiber.Map{
		"Title": "Single Task",
		"Task":  task,
	})
}

func NotFound(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).SendFile("./public/404.html")
}

func EditTask(c *fiber.Ctx) error {
	task := models.Task{}
	id := c.Params("id")

	result := database.DB.Db.Where("id = ?", id).First(&task)
	if result.Error != nil {
		return NotFound(c)
	}

	return c.Render("edit", fiber.Map{
		"Title":    "Edit Task",
		"Subtitle": "Edit your task",
		"Task":     task,
	})
}

func UpdateTask(c *fiber.Ctx) error {
	task := new(models.Task)
	id := c.Params("id")

	if err := c.BodyParser(task); err != nil {
		return c.Status(fiber.StatusServiceUnavailable).SendString(err.Error())
	}

	result := database.DB.Db.Model(&task).Where("id = ?", id).Updates(task)
	if result.Error != nil {
		return EditTask(c)
	}

	return ShowTask(c)
}

func DeleteTask(c *fiber.Ctx) error {
	task := models.Task{}
	id := c.Params("id")

	result := database.DB.Db.Where("id = ?", id).Delete(&task)
	if result.Error != nil {
		return NotFound(c)
	}

	return ListTasks(c)
}
