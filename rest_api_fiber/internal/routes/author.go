package routes

import (
	"github.com/andreichristian/rest/internal/db"
	"github.com/andreichristian/rest/internal/handlers"
	"github.com/andreichristian/rest/internal/models"
	"github.com/gofiber/fiber/v2"
)

func AuthorRoutes(r fiber.Router) {

	r.Get("/", handlers.GetAllAuthors)

	r.Get("/:authorId", handlers.GetOneAuthor)

	r.Post("/", handlers.PostAuthor )

	r.Patch("/:authorId",handlers.PatchAuthor)

	r.Delete("/:authorId", func(c *fiber.Ctx) error {

		authorId := c.Params("authorId")

		result := db.DB.Delete(&models.Author{}, authorId)

		if result.Error != nil {

			return result.Error

		}

		return c.JSON(fiber.Map{"ok": true})
	})



}
