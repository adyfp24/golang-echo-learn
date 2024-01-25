package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "Halo ady, server is running on port 8000",
		})
	})
	err := app.Listen(":8000")
	if err != nil {
		log.Fatal(err)
	}
}
