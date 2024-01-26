package main

import (
	"github.com/adyfp24/golang-fiber-learn/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	routes.RouteInit(app)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("welcome ady, server is running on port 8000")
	})
	err := app.Listen(":8000")
	if err != nil {
		log.Fatal(err)
	}
}
