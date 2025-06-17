package main

import (
	"github.com/labstack/echo/v4"
	bookHand "go_library/internal/api/handler/book"
	"go_library/internal/api/routers"
	bookServ "go_library/internal/domain/book"
	"go_library/internal/infrastructure/db"
	bookRepo "go_library/internal/infrastructure/repository/book"
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
	bookRepository := bookRepo.NewBookRepository(database)
	bookServise := bookServ.NewBookService(bookRepository)
	bookHadler := bookHand.NewBookHandler(bookServise)
	routers.RegisterRoutes(e, bookHadler)
	e.Start("localhost:8000")
}
