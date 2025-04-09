package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"github.com/dinukaDB/fiber-books-api/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&models.Book{})
	DB = database
}
