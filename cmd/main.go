package main

import (
	"github.com/labstack/echo/v4"
	"go_library/handlers"
	"go_library/internal/infrastructure/db"
	"go_library/repository"
	"go_library/service"
	"log"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	e := echo.New()
	//bookRepo := repository.NewBookRepository(database)
	//bookService := service.NewBookService(bookRepo)
	//bookHadler := handlers.NewBookHandler(bookService)
	//e.GET("/books", bookHadler.GetAllBooks)
	//e.GET("/books/:id", bookHadler.GetBookByID)
	//e.POST("/books", bookHadler.CreateBook)
	//e.PATCH("/books/:id", bookHadler.UpdateBook)
	//e.DELETE("/books/:id", bookHadler.DeleteBook)
	e.Start("localhost:8000")
}
