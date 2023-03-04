package routes

import (
	"github.com/gofiber/fiber/v2"
	bookController "github.com/guisantosalves/go-api-fiber/src/controllers"
	"github.com/guisantosalves/go-api-fiber/src/database"
	"github.com/guisantosalves/go-api-fiber/src/models"
	"gorm.io/gorm/logger"
)

func SetRoutes(app *fiber.App) {
	v1 := app.Group("/api/v1")

	v1.Get("/book", bookController.GetBooks)
	v1.Get("/book/:id", bookController.GetBookById)
	v1.Post("/book", bookController.NewBook)
	v1.Delete("/book/:id", bookController.DeleteBook)
}
