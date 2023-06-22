package handlers

import (
	"fiber/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetBooks(c *fiber.Ctx) error {

	book := models.Book{
		Id:     1,
		Author: "Fitzgerald",
		Title:  "Great Gatsby",
		Year:   1920,
	}

	return c.JSON(book)
}

func PostBooks(c *fiber.Ctx) error {

	book := new(models.Book)

	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error()})
	}

	fmt.Println(book)

	return c.JSON(book)
}
