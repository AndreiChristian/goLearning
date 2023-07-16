package routes

import "github.com/gofiber/fiber/v2"

func BookRoutes(r fiber.Router) {

	r.Get("/", func(c *fiber.Ctx) error {

		return c.JSON(fiber.Map{"hello": "world"})

	})

}
