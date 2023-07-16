package main

import (
	"fmt"
	"log"

	"github.com/andreichristian/rest/internal/db"
	"github.com/andreichristian/rest/internal/routes"
	"github.com/andreichristian/rest/migrations"
	"github.com/gofiber/fiber/v2"
)

func init() {

	err := db.NewDB()

	if err != nil {

		log.Fatalf("Could not connect to the db :%v \n", err)

	}

	fmt.Println("Connected to the db")

	migrations.Migrate()

	fmt.Println("Migrated")

}

func main() {

	app := fiber.New()

	api := app.Group("/api")

	routes.BookRoutes(api.Group("/books"))

	routes.AuthorRoutes(api.Group("/authors"))

	log.Fatal(app.Listen(":3000"))

}
