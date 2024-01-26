package main

import (
	"log"
	"github.com/adyfp24/golang-fiber-learn/database/migrations"
	"github.com/adyfp24/golang-fiber-learn/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	migrations.RunMigration()

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
