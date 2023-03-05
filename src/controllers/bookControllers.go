package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/guisantosalves/go-api-fiber/src/database"
	"github.com/guisantosalves/go-api-fiber/src/models"
	"gorm.io/gorm/logger"
)

func GetBooks(c *fiber.Ctx) error { // array of Book
	var books []models.Book
	dta := database.DB.Db

	if result := dta.Find(&books); result.Error != nil {
		return c.Status(fiber.StatusBadRequest).SendString(result.Error.Error())
	}

	return c.Status(fiber.StatusAccepted).JSON(&books)
}

func GetBookById(c *fiber.Ctx) error {
	idFromParams := c.Params("id")
	dta := database.DB.Db
	var book models.Book

	if result := dta.First(&book, idFromParams); result.Error != nil {
		return c.Status(fiber.StatusBadRequest).SendString(result.Error.Error())
	}

	return c.Status(fiber.StatusAccepted).JSON(book)
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

func UpdateBookById(c *fiber.Ctx) error {
	var bookFromBody models.Book
	dta := database.DB.Db
	idFromParams := c.Params("id")
	var bookFromDbToUpdate models.Book

	// mapped everything from body to bookFromBody
	if err := c.BodyParser(&bookFromBody); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// insert into database - get the data that I want to update and save with the
	// new data from body
	if result := dta.First(&bookFromDbToUpdate, idFromParams); result.Error != nil {
		return c.Status(fiber.StatusBadRequest).SendString(result.Error.Error())
	}

	bookFromDbToUpdate.Title = bookFromBody.Title
	bookFromDbToUpdate.Author = bookFromBody.Author
	bookFromDbToUpdate.Desc = bookFromBody.Desc

	// saving the changes
	dta.Save(&bookFromDbToUpdate)

	return c.Status(fiber.StatusAccepted).JSON(&bookFromDbToUpdate)
}

func DeleteBookById(c *fiber.Ctx) error {
	idFromParams := c.Params("id")
	var book models.Book
	dta := database.DB.Db

	if result := dta.Delete(&book, idFromParams); result.Error != nil {
		return c.Status(fiber.StatusBadRequest).SendString(result.Error.Error())
	}

	return c.Status(fiber.StatusAccepted).JSON(book)
}
