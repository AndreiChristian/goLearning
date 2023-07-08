package controller

import "github.com/gofiber/fiber/v2"

func GetAllBooks(c *fiber.Ctx) error {

	return c.SendString("All books routes \n")

}
