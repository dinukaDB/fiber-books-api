package main

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/dinukaDB/fiber-books-api/handlers"
    "github.com/dinukaDB/fiber-books-api/database"
    "github.com/gofiber/fiber/v2"
    "github.com/stretchr/testify/assert"
)

func setupApp() *fiber.App {
    app := fiber.New()
    database.ConnectDatabase()
    setupRoutes(app)
    return app
}

func setupRoutes(app *fiber.App) {
    app.Post("/books", handlers.CreateBook)
    app.Get("/books", handlers.GetBooks)
    app.Get("/books/:id", handlers.GetBook)
    app.Put("/books/:id", handlers.UpdateBook)
    app.Delete("/books/:id", handlers.DeleteBook)
}

func TestCreateBook(t *testing.T) {
    app := setupApp()

    payload := map[string]interface{}{
        "title":  "Test Book",
        "author": "Test Author",
        "year":   2025,
    }
    body, _ := json.Marshal(payload)

    req := httptest.NewRequest("POST", "/books", bytes.NewReader(body))
    req.Header.Set("Content-Type", "application/json")

    resp, err := app.Test(req)
    assert.NoError(t, err)
    assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestGetBooks(t *testing.T) {
    app := setupApp()

    req := httptest.NewRequest("GET", "/books", nil)
    resp, err := app.Test(req)

    assert.NoError(t, err)
    assert.Equal(t, http.StatusOK, resp.StatusCode)
}
