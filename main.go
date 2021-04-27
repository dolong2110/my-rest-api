package main

import (
	"github.com/gofiber/fiber"
	"my-rest-api/common"
	"my-rest-api/handlers"
)

func main() {
	app := fiber.New()

	app.Get("/person/:id?", handlers.GetPerson)
	app.Post("/person", handlers.CreatePerson)
	app.Put("/person/:id", handlers.UpdatePerson)
	app.Delete("/person/:id", handlers.DeletePerson)

	app.Listen(common.Port)
}
