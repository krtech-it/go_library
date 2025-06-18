package routers

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	authHand "go_library/internal/api/handler/auth"
	authorHand "go_library/internal/api/handler/author"
	bookHand "go_library/internal/api/handler/book"
)

var JwtSecret = []byte("string")

func RegisterRoutes(e *echo.Echo,
	bookHandler *bookHand.BookHandler,
	authorHandler *authorHand.AuthorHandler,
	authHandler *authHand.AuthHandler) {

	e.POST("/login", authHandler.Login)

	api := e.Group("/api")
	api.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: JwtSecret,
	}))

	bookGroup := api.Group("/book")
	bookGroup.GET("", bookHandler.GetAllBooks)
	bookGroup.GET("/:id", bookHandler.GetBookByID)
	bookGroup.POST("", bookHandler.CreateBook)
	bookGroup.PATCH("/:id", bookHandler.UpdateBook)
	bookGroup.DELETE("/:id", bookHandler.DeleteBook)

	authorGroup := api.Group("/author")
	authorGroup.GET("", authorHandler.GetAllAuthors)
	authorGroup.GET("/:id", authorHandler.GetAuthorByID)
}
