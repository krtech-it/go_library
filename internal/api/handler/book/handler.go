package book

import (
	"github.com/labstack/echo/v4"
	"go_library/internal/api/dto"
	"go_library/internal/domain/book"
	"go_library/internal/utils/mapper"
	"net/http"
)

type BookHandler struct {
	service book.BookService
}

func (h *BookHandler) GetAllBooks(c echo.Context) error {
	books, err := h.service.GetAllBooks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not get books"})
	}
	response := make([]*dto.BookResponse, 0)
	for _, value := range books {
		response = append(response, mapper.ToBookResponse(value))
	}
	return c.JSON(http.StatusOK, response)
}

func (h *BookHandler) GetBookByID(c echo.Context) error {
	id := c.Param("id")
	book, err := h.service.GetBookByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not get book"})
	}
	result := mapper.ToBookResponse(book)
	return c.JSON(http.StatusOK, result)
}

func NewBookHandler(service book.BookService) *BookHandler {
	return &BookHandler{service: service}
}
