package routers

import (
	"github.com/labstack/echo/v4"
	authorHand "go_library/internal/api/handler/author"
	bookHand "go_library/internal/api/handler/book"
)

func RegisterRoutes(e *echo.Echo,
	bookHandler *bookHand.BookHandler,
	authorHandler *authorHand.AuthorHandler) {
	api := e.Group("/api")

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
