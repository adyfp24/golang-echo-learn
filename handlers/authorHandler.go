package handlers

import (
	"errors"
	"github.com/adyfp24/golang-fiber-learn/database"
	"github.com/adyfp24/golang-fiber-learn/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateAuthor(c *fiber.Ctx) error {
	db := database.DB
	author := new(models.Author)
	err := c.BodyParser(author)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "bad gateway",
			"error":   err.Error(),
		})
	}
	result := db.Create(&author).Error
	if result != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "gagal menambahkan data",
			"error":   err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"succes":  "true",
		"message": "berhasil menambahkan data author",
		"data":    author,
	})
}

func GetAllAuthor(c *fiber.Ctx) error {
	db := database.DB
	var allAuthor []models.Author
	err := db.Find(&allAuthor).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "bad gateway",
			"error":   err.Error(),
		})
	}
	if len(allAuthor) == 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "data tidak ditemukan",
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"succes":  "true",
		"message": "berhasil menambahkan data semua author",
		"data":    allAuthor,
	})
}

func GetAuthorById(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("authorId")
	var author models.Author
	err := db.Find(&author, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{
				"message": "id tidak ditemukan",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"message": "bad gateway",
			"error":   err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"succes":  "true",
		"message": "berhasil mendapat data author",
		"data":    author,
	})
}
func UpdateAuthor(c *fiber.Ctx) error {
	return nil
}
func DeleteAuthor(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("authorId")
	var deletedAuthor models.Author
	err := db.First(&deletedAuthor, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{
				"message": "id tidak ditemukan",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"message": "bad gateway",
			"error":   err.Error(),
		})
	}
	err = db.Delete(&deletedAuthor, id).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "gagal menghapus data author",
			"error":   err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"succes":  "true",
		"message": "berhasil menghapus data author",
		"data": deletedAuthor,
	})
}
