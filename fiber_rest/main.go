package main

import (
	"fiber_rest/api/router"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Greeting struct {
	Message string `json:"name"`
}

func main() {

	app := fiber.New()

	api := app.Group("/api")

	router.Router(api)

	log.Fatal(app.Listen(":3000"))

}
