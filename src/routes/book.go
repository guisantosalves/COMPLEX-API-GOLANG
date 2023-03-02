package routes

import (
	"github.com/gofiber/fiber"
	bookController "github.com/guisantosalves/go-api-fiber/src/controllers"
	"github.com/guisantosalves/go-api-fiber/src/database"
	"github.com/guisantosalves/go-api-fiber/src/models"
)

func SetRoutes(app *fiber.App) {
	v1 := app.Group("/api/v1", func(c *fiber.Ctx) {
		c.Set("Version", "v1")
	})

	v1.Get("/book", bookController.GetBooks)
	v1.Get("/book/:id", bookController.GetBookById)
	v1.Post("/book", func(c *fiber.Ctx) {
		// like an instance
		body := models.Book{}

		// parsing body and mapping it into models.Book struct
		if err := c.BodyParser(&body); err != nil {
			// handling with error
			c.Status(fiber.StatusBadRequest).Send(err.Error())
		}

		// make book receive all mapped data from body
		var book models.Book
		book.Title = body.Title
		book.Desc = body.Desc
		book.Author = body.Author

		//insert db
		if result := database.DB.Create(&book); result.Error != nil {
			panic(fiber.ErrInternalServerError)
		}

		c.Status(fiber.StatusCreated).JSON(&book)
	})
	v1.Delete("/book/:id", bookController.DeleteBook)
}
