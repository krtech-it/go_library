package handlers

import (
	"github.com/labstack/echo/v4"
	"go_library/schemas"
	"go_library/service"
	"net/http"
)

type BookHandler struct {
	service service.BookService
}

func (h *BookHandler) GetAllBooks(c echo.Context) error {
	result, err := h.service.GetAllBooks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not get books"})
	}
	return c.JSON(http.StatusOK, result)
}

func (h *BookHandler) GetBookByID(c echo.Context) error {
	id := c.Param("id")
	result, err := h.service.GetBookByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not get book"})
	}
	return c.JSON(http.StatusOK, result)
}

func (h *BookHandler) CreateBook(c echo.Context) error {
	var req schemas.BookRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}
	bookId, err := h.service.CreateBook(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not get book"})
	}
	return c.JSON(http.StatusCreated, bookId)
}

func (h *BookHandler) UpdateBook(c echo.Context) error {
	var req schemas.BookRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}
	id := c.Param("id")
	bookId, err := h.service.UpdateBook(id, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not get book"})
	}
	return c.JSON(http.StatusNoContent, bookId)
}

func (h *BookHandler) DeleteBook(c echo.Context) error {
	id := c.Param("id")
	err := h.service.DeleteBook(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not get book"})
	}
	return c.NoContent(http.StatusNoContent)

}

func NewBookHandler(service service.BookService) *BookHandler {
	return &BookHandler{service: service}
}
