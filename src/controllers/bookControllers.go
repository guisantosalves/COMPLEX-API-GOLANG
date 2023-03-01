package controllers

import (
	"github.com/gofiber/fiber"
)

func GetBooks(c *fiber.Ctx) {
	c.Send("getBooks")
}

func GetBookById(c *fiber.Ctx) {
	c.Send("getBooksbyid")
}

func NewBook(c *fiber.Ctx) {
	c.Send("create book")
}

func DeleteBook(c *fiber.Ctx) {
	c.Send("delete book")
}
