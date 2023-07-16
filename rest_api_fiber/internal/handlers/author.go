package handlers

import (
	"github.com/andreichristian/rest/internal/db"
	"github.com/andreichristian/rest/internal/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllAuthors(c *fiber.Ctx) error {

	var authors []models.Author

	result := db.DB.Find(&authors)

	if result.Error != nil {

		return result.Error

	}

	return c.JSON(authors)

}

func GetOneAuthor(c *fiber.Ctx) error {

	authorId := c.Params("authorId")

	var author models.Author

	result := db.DB.First(&author, authorId)

	if result.Error != nil {

		return result.Error

	}

	return c.JSON(author)

}

func PostAuthor(c *fiber.Ctx) error {

		author := &models.Author{}

		err := c.BodyParser(author)
		if err != nil {
			return err
		}

		result := db.DB.Create(author)
		if result.Error != nil {
			return result.Error
		}

		return c.JSON(author)

	}

func PatchAuthor(c *fiber.Ctx ) error{

	authorId := c.Params("authorId")

	var author models.Author

	result := db.DB.First(&author, authorId)

	if result.Error !=nil {
		
		return result.Error
	}

	err := c.BodyParser(&author)
	if err!=nil{

		return err
	}

	db.DB.Save(&author)

	return c.JSON(author)

}
