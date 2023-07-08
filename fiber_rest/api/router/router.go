package router

import "github.com/gofiber/fiber/v2"

func Router(router fiber.Router) {

	router.Get("/users", func(c *fiber.Ctx) error {

		return c.Status(fiber.StatusOK).SendString("Users route \n")

	})

	router.Get("/user/:userId", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("User " + c.Params("userId") + "\n")
	})

}
