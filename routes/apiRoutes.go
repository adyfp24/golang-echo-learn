package routes

import(
	"github.com/gofiber/fiber/v2"
	"github.com/adyfp24/golang-fiber-learn/handlers"
)

func RouteInit(app *fiber.App){
	r := app.Group("/api")

	r.Post("/author", handlers.CreateAuthor)

	r.Get("/author", handlers.GetAllAuthor)

	r.Get("author/:idAuthor", handlers.GetAuthorById)

	r.Put("author/:idAuthor", handlers.UpdateAuthor)

	r.Delete("author/:idAuthor", handlers.DeleteAuthor)
}

