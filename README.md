# 📚 Book Management API - Go + Fiber
A simple RESTful API for managing a collection of books built using **GoLang (v1.18+)**, **Fiber v2**, and **SQLite**. This project is part of an internship test assignment.

---
## 🚀 Features
- Create, Read, Update, and Delete (CRUD) books
- Clean modular structure (Handlers, Services, Models, DB)
- Unit testing for handlers
- SQLite integration using GORM
- Ready-to-use REST endpoints

---
## 🔧 Tech Stack
- **GoLang v1.18+**
- **Fiber v2** - Web framework
- **GORM** - ORM for SQLite
- **SQLite** - Local database
- **Testify** - For unit testing

---
## 📘 Book Model
```go
type Book struct {
    ID     uint   `json:"id" gorm:"primaryKey"`
    Title  string `json:"title"`
    Author string `json:"author"`
    Year   int    `json:"year"`
}
```
# Running the Project
## Clone the repo
git clone https://github.com/dinukaDB/fiber-books-api.git
cd fiber-books-api

## Install dependencies
go mod tidy

## Run the app
go run main.go

# Testing the API
## Using Curl
### Create Book
```
curl -X POST http://localhost:3000/books -H "Content-Type: application/json" -d "{\"title\":\"1984\", \"author\":\"George Orwell\", \"year\":1949}"
```

### Get all books
```
curl http://localhost:3000/books
```

### Get a book by ID
```
curl http://localhost:3000/books/1
```

### Update a book
```
curl -X PUT http://localhost:3000/books/1 -H "Content-Type: application/json" -d "{\"title\":\"Updated\", \"author\":\"Someone\", \"year\":2025}"
```

### Delete a book
```
curl -X DELETE http://localhost:3000/books/1
```
![image1] (media/image1.png)

