package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/dinukaDB/fiber-books-api/models"
	"github.com/dinukaDB/fiber-books-api/services"
)

func GetBooks(c *fiber.Ctx) error {
	books := services.GetAllBooks()
	return c.JSON(books)
}

func GetBook(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	book, err := services.GetBookByID(uint(id))
	if err != nil {
		return c.Status(404).SendString("Book not found")
	}
	return c.JSON(book)
}

func CreateBook(c *fiber.Ctx) error {
	var book models.Book
	if err := c.BodyParser(&book); err != nil {
		return c.Status(400).SendString("Invalid request")
	}
	if err := services.CreateBook(&book); err != nil {
		return c.Status(500).SendString("Could not create book")
	}
	return c.JSON(book)
}

func UpdateBook(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	book, err := services.GetBookByID(uint(id))
	if err != nil {
		return c.Status(404).SendString("Book not found")
	}
	if err := c.BodyParser(&book); err != nil {
		return c.Status(400).SendString("Invalid request")
	}
	if err := services.UpdateBook(&book); err != nil {
		return c.Status(500).SendString("Update failed")
	}
	return c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if err := services.DeleteBook(uint(id)); err != nil {
		return c.Status(500).SendString("Delete failed")
	}
	return c.SendStatus(204)
}
