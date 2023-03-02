package controllers

import (
	"github.com/gofiber/fiber"
	"github.com/guisantosalves/go-api-fiber/src/database"
	"github.com/guisantosalves/go-api-fiber/src/models"
)

func GetBooks(c *fiber.Ctx) {
	c.Send("getBooks")
}

func GetBookById(c *fiber.Ctx) {
	c.Send("getBooksbyid")
}

func NewBook(c *fiber.Ctx) error {
	// like an instance
	body := models.Book{}

	// parsing body and mapping it into models.Book struct
	if err := c.BodyParser(&body); err != nil {
		// handling with error
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// make book receive all mapped data from body
	var book models.Book
	book.Title = body.Title
	book.Desc = body.Desc
	book.Author = body.Author

	//insert db
	if result := database.DB.Create(&book); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&book)
}

func DeleteBook(c *fiber.Ctx) {
	c.Send("delete book")
}
