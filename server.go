package main

import (
	"app/controller"
	"app/database"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	App := fiber.New(fiber.Config{
		Prefork: true,
	})

	database.DbConnection()
	database.SyncDb()

	App.Post("/todo", controller.Create) //This Post Method Create A Todo Note
	App.Get("/todo/:id", controller.GetById)
	App.Get("/todo", controller.GetAll)
	App.Delete("/todo/:id", controller.Delete)

	log.Fatal(App.Listen(":3000"))

}
