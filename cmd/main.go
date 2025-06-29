package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "go_library/docs"
	authHand "go_library/internal/api/handler/auth"
	authorHand "go_library/internal/api/handler/author"
	bookHand "go_library/internal/api/handler/book"
	"go_library/internal/api/routers"
	authServ "go_library/internal/domain/auth"
	authorServ "go_library/internal/domain/author"
	bookServ "go_library/internal/domain/book"
	"go_library/internal/infrastructure/db"
	middleware2 "go_library/internal/infrastructure/middleware_custom"
	authRepo "go_library/internal/infrastructure/repository/auth"
	authorRepo "go_library/internal/infrastructure/repository/author"
	bookRepo "go_library/internal/infrastructure/repository/book"
	"log"
)

// @title My API
// @version 1.0
// @description Это Swagger API для Go проекта
// @host localhost:8000
// @BasePath /
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware2.ErrorMiddleware)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Auth init
	authRepository := authRepo.NewAuthRepository(database)
	authService := authServ.NewAuthService(authRepository)
	authHand := authHand.NewAuthHandler(authService)

	// Book init
	bookRepository := bookRepo.NewBookRepository(database)
	bookService := bookServ.NewBookService(bookRepository)
	bookHandler := bookHand.NewBookHandler(bookService)

	// Author init
	authorRepository := authorRepo.NewAuthorRepository(database)
	authorService := authorServ.NewAuthorService(authorRepository)
	authorHandler := authorHand.NewAuthorHandler(authorService)

	routers.RegisterRoutes(e, bookHandler, authorHandler, authHand)
	e.Start("localhost:8000")
}
