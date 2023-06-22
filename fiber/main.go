package main

import (
	"fiber/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	api := app.Group("/api")

	api.Get("/books", handlers.GetBooks)
	api.Post("/books", handlers.PostBooks)

	app.Listen(":3000")
}
