package handlers

import (
	"github.com/adyfp24/golang-fiber-learn/models"
	"github.com/gofiber/fiber/v2"
)

func CreateAuthor(c *fiber.Ctx) error {
	author := new(models.Author)
	err := c.BodyParser(author)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "bad gateway",
			"error":   err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"succes": "true",
		"message":  "berhasil menambahkan data author",
	})
}

func GetAllAuthor(c *fiber.Ctx) error {
	return nil
}

func GetAuthorById(c *fiber.Ctx) error {
	return nil
}
func UpdateAuthor(c *fiber.Ctx) error {
	return nil
}
func DeleteAuthor(c *fiber.Ctx) error {
	return nil
}
