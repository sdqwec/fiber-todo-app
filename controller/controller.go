package controller

import (
	"app/database"
	"app/models"

	"github.com/gofiber/fiber/v2"
)

func Create(c *fiber.Ctx) error {
	db := database.DBConn
	todo := new(models.ToDo)

	if err := c.BodyParser(todo); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&todo)
	return c.Status(200).JSON(todo)
}

func GetById(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")

	var todo models.ToDo
	db.Find(&todo, id)
	return c.Status(200).JSON(todo)

}

func GetAll(c *fiber.Ctx) error {
	db := database.DBConn
	var todos []models.ToDo
	db.Find(&todos)
	return c.Status(200).JSON(todos)
}

func Delete(c *fiber.Ctx) error {
	db := database.DBConn
	var todo models.ToDo
	id := c.Params("id")

	db.First(&todo, id)
	if todo.Name == "" {
		return c.Status(500).SendString("Book not Found")
	}
	db.Delete(&todo)
	return c.Status(200).JSON("Deleted")
}
