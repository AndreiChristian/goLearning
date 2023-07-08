package router

import (
	"fiber_rest/api/controller"

	"github.com/gofiber/fiber/v2"
)

func BookRouter(router fiber.Router) {

	router.Get("/", controller.GetAllBooks)

	router.Get("/:bookId", func(c *fiber.Ctx) error {
		return c.SendString("Book " + c.Params("bookId") + "/")
	})

}
