package book

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"go_library/internal/api/dto"
	"go_library/internal/domain/book"
	ApiError "go_library/internal/errors"
	"go_library/internal/utils/mapper"
	"net/http"
	"strconv"
)

type BookHandler struct {
	service book.BookService
}

// GetAllBooks godoc
// @Summary Получить все книги
// @Description Возвращает список всех книг с автором и мета-данными
// @Tags books
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {array} dto.BookResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/book [get]
func (h *BookHandler) GetAllBooks(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		return ApiError.NewAPIError(http.StatusBadRequest, "page param is invalid")
	}
	pageSize, err := strconv.Atoi(c.QueryParam("size"))
	if err != nil {
		return ApiError.NewAPIError(http.StatusBadRequest, "size param is invalid")
	}
	books, count, err := h.service.GetAllBooks(page, pageSize)
	if err != nil {
		return err
	}
	booksDto := make([]*dto.BookResponse, 0)
	for _, value := range books {
		booksDto = append(booksDto, mapper.ToBookResponse(value))
	}
	response := &dto.BookResponsePagination{
		PageStruct: dto.PageStruct{
			Page:     page,
			PageSize: pageSize,
			Count:    count,
		},
		Books: booksDto,
	}
	return c.JSON(http.StatusOK, response)
}

// GetBookByID godoc
// @Summary Получить книгу по ID
// @Description Возвращает книгу по её идентификатору с информацией об авторе
// @Tags books
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Book ID" example("123e4567-e89b-12d3-a456-426614174000")
// @Success 200 {object} dto.BookResponseGenres
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/book/{id} [get]
func (h *BookHandler) GetBookByID(c echo.Context) error {
	id := c.Param("id")
	bookObj, err := h.service.GetBookByID(id)
	if err != nil {
		return err
	}
	result := mapper.ToBookResponseWithGenres(bookObj)
	return c.JSON(http.StatusOK, result)
}

// CreateBook godoc
// @Summary Создать новую книгу
// @Description Создает новую книгу в библиотеке с указанным автором
// @Tags books
// @Accept json
// @Produce json
// @Security Bearer
// @Param book body dto.BookRequest true "Book object"
// @Success 201 {object} dto.BookIdResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/book [post]
func (h *BookHandler) CreateBook(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)
	var req dto.BookRequest
	if err := c.Bind(&req); err != nil {
		return ApiError.NewAPIError(http.StatusBadRequest, "internal server error")
	}
	domainBook := mapper.FromRequestToDomainBook(&req)
	bookId, err := h.service.CreateBook(domainBook, userID)
	if err != nil {
		return err
	}
	response := dto.BookIdResponse{
		ID: bookId,
	}
	return c.JSON(http.StatusCreated, response)
}

// UpdateBook godoc
// @Summary Обновить книгу
// @Description Обновляет информацию о книге по её ID
// @Tags books
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Book ID" example("123e4567-e89b-12d3-a456-426614174000")
// @Param book body dto.BookRequest true "Book object"
// @Success 202 {object} dto.BookIdResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/book/{id} [patch]
func (h *BookHandler) UpdateBook(c echo.Context) error {
	var req dto.BookRequest
	if err := c.Bind(&req); err != nil {
		return ApiError.NewAPIError(http.StatusBadRequest, "internal server error")
	}
	id := c.Param("id")
	domainBook := mapper.FromRequestToDomainBook(&req)
	bookId, err := h.service.UpdateBook(id, domainBook)
	if err != nil {
		return err
	}
	response := dto.BookIdResponse{
		ID: bookId,
	}
	return c.JSON(http.StatusAccepted, response)
}

// DeleteBook godoc
// @Summary Удалить книгу
// @Description Удаляет книгу по её ID
// @Tags books
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Book ID" example("123e4567-e89b-12d3-a456-426614174000")
// @Success 204 "No Content"
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/book/{id} [delete]
func (h *BookHandler) DeleteBook(c echo.Context) error {
	id := c.Param("id")
	err := h.service.DeleteBook(id)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}

func NewBookHandler(service book.BookService) *BookHandler {
	return &BookHandler{service: service}
}
