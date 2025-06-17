package routers

import (
	"github.com/labstack/echo/v4"
	bookHandler "go_library/internal/api/handler/book"
)

func RegisterRoutes(e *echo.Echo, bookHandler *bookHandler.BookHandler) {
	api := e.Group("/api")

	bookGroup := api.Group("/book")
	bookGroup.GET("/", bookHandler.GetAllBooks)
	bookGroup.GET("/:id", bookHandler.GetBookByID)
	bookGroup.POST("/", bookHandler.CreateBook)
}
