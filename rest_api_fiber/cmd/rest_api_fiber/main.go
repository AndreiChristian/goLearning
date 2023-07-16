package main

import (
	"fmt"
	"log"

	"github.com/andreichristian/rest/internal/db"
	"github.com/gofiber/fiber/v2"
)

func init() {

	err := db.NewDB()

	if err != nil {

		log.Fatalf("Could not connect to the db :%v \n", err)

	}

	fmt.Println("Connected to the db")

}

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {

		return c.SendString("Hello World")
	})

	log.Fatal(app.Listen(":3000"))

}
