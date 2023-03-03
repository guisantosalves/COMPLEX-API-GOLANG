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
	v1.Post("/book", func(c *fiber.Ctx) error {
		// like an instance
		body := models.Book{}
		dta := database.DB.Db
		// parsing body and mapping it into models.Book struct
		if err := c.BodyParser(&body); err != nil {
			// handling with error
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		// make book receive all mapped data from body
		var book models.Book
		book.Title = body.Title
		book.Desc = body.Desc
		book.Author = body.Author

		//insert db
		if result := dta.Create(&book); result.Error != nil {
			logger.Default.LogMode(logger.Info)
			panic(fiber.ErrInternalServerError)
		}

		return c.Status(fiber.StatusCreated).JSON(&book)
	})
	v1.Delete("/book/:id", bookController.DeleteBook)
}
