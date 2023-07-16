package routes

import (
	"github.com/andreichristian/rest/internal/models"
	"github.com/gofiber/fiber/v2"
)

func BookRoutes(r fiber.Router) {

	r.Get("/", func(c *fiber.Ctx) error {

		return c.JSON(fiber.Map{"hello": "world"})

	})

	r.Post("/", func(c *fiber.Ctx) error {

		a:= new(models.Author)

		err := c.BodyParser(a);

		if err!=nil{
	
			return err
		}

		return c.JSON(a)

	})


}
