package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/guisantosalves/go-api-fiber/src/database"
	"github.com/guisantosalves/go-api-fiber/src/models"
)

func GetBooks(c *fiber.Ctx) error {
	return c.SendString("getBooks")
}

func GetBookById(c *fiber.Ctx) error {
	// array of Book
	var books []models.Book{}
	dta := database.DB.db

	if result := dta.Find(&books); result.Error != nil {
		return c.Status(fiber.StatusBadRequest).SendString(result.Error());
	}

	return c.Status(fiber)
}

func NewBook(c *fiber.Ctx) error {
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
}

func DeleteBook(c *fiber.Ctx) error {
	return c.SendString("delete book")
}
