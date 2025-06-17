package book

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go_library/internal/api/dto"
	"go_library/internal/domain/book"
	"go_library/internal/utils/mapper"
	"net/http"
)

type BookHandler struct {
	service book.BookService
}

func (h *BookHandler) GetAllBooks(c echo.Context) error {
	log.Info("Get all books")
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

func (h *BookHandler) CreateBook(c echo.Context) error {
	var req dto.BookRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}
	domainBook := mapper.FromRequestToDomain(&req)
	bookId, err := h.service.CreateBook(domainBook)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not get book"})
	}
	response := dto.BookIdResponse{
		ID: bookId,
	}
	return c.JSON(http.StatusCreated, response)
}

func (h *BookHandler) UpdateBook(c echo.Context) error {
	var req dto.BookRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}
	id := c.Param("id")
	domainBook := mapper.FromRequestToDomain(&req)
	bookId, err := h.service.UpdateBook(id, domainBook)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not get book"})
	}
	response := dto.BookIdResponse{
		ID: bookId,
	}
	return c.JSON(http.StatusAccepted, response)
}

func (h *BookHandler) DeleteBook(c echo.Context) error {
	id := c.Param("id")
	err := h.service.DeleteBook(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not get book"})
	}
	return c.NoContent(http.StatusNoContent)
}

func NewBookHandler(service book.BookService) *BookHandler {
	return &BookHandler{service: service}
}
