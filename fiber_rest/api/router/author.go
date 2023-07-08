package router

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func AuthorsRouter(router fiber.Router) {

	router.Use(func(c *fiber.Ctx) error {

		log.Println("Used the Authors Router")

		return c.Next()

	})

	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("All author routes \n")
	})

	router.Get("/:authorId", func(c *fiber.Ctx) error {
		return c.SendString("Author " + c.Params("authorId") + "/")
	})

}
