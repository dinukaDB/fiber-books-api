package services

import (
	"github.com/dinukaDB/fiber-books-api/models"
	"github.com/dinukaDB/fiber-books-api/database"
)

func GetAllBooks() []models.Book {
	var books []models.Book
	database.DB.Find(&books)
	return books
}

func GetBookByID(id uint) (models.Book, error) {
	var book models.Book
	result := database.DB.First(&book, id)
	return book, result.Error
}

func CreateBook(book *models.Book) error {
	return database.DB.Create(book).Error
}

func UpdateBook(book *models.Book) error {
	return database.DB.Save(book).Error
}

func DeleteBook(id uint) error {
	return database.DB.Delete(&models.Book{}, id).Error
}
