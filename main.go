package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/dinukaDB/fiber-books-api/database"
	"github.com/dinukaDB/fiber-books-api/handlers"
)

func main() {
	app := fiber.New()
	database.ConnectDatabase()

	app.Get("/books", handlers.GetBooks)
	app.Get("/books/:id", handlers.GetBook)
	app.Post("/books", handlers.CreateBook)
	app.Put("/books/:id", handlers.UpdateBook)
	app.Delete("/books/:id", handlers.DeleteBook)

	app.Listen(":3000")
}
