package routes

import (
	"github.com/gofiber/fiber"
	bookController "github.com/guisantosalves/go-api-fiber/src/controllers"
)

func SetRoutes(app *fiber.App) {
	v1 := app.Group("/api/v1")
	v1.Get("/book", bookController.GetBooks)
	v1.Get("/book/:id", bookController.GetBookById)
	v1.Post("/book", bookController.NewBook)
	v1.Delete("/book/:id", bookController.DeleteBook)
}
