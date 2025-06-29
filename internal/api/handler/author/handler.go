package author

import (
	"github.com/labstack/echo/v4"
	"go_library/internal/api/dto"
	"go_library/internal/api/handler"
	"go_library/internal/api/mapper/fromDto"
	"go_library/internal/api/mapper/toDto"
	"go_library/internal/domain/author"
	ApiError "go_library/internal/errors"
	"net/http"
)

type AuthorHandler struct {
	service author.AuthorService
}

// GetAllAuthors godoc
// @Summary Получить всех авторов
// @Description Возвращает список всех авторов
// @Tags authors
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {array} dto.AuthorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/author [get]
func (h *AuthorHandler) GetAllAuthors(c echo.Context) error {
	authors, err := h.service.GetAllAuthors()
	if err != nil {
		return err
	}
	response := make([]*dto.AuthorResponse, 0)
	for _, value := range authors {
		response = append(response, toDto.ToDtoAuthor(value))
	}
	return c.JSON(http.StatusOK, response)
}

// GetAuthorByID godoc
// @Summary Получить автора по ID
// @Description Возвращает информацию об авторе по его идентификатору, включая список его книг
// @Tags authors
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Author ID" example("123e4567-e89b-12d3-a456-426614174000")
// @Success 200 {object} dto.AuthorFullResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/author/{id} [get]
func (h *AuthorHandler) GetAuthorByID(c echo.Context) error {
	id := c.Param("id")
	authorObj, err := h.service.GetAuthorByID(id)
	if err != nil {
		return err
	}
	response := toDto.ToDtoAuthorWithBooks(authorObj)
	return c.JSON(http.StatusOK, response)
}

// CreateAuthor godoc
// @Summary Создать автора
// @Description Создает нового автора. Требуется авторизация.
// @Tags authors
// @Accept json
// @Produce json
// @Security Bearer
// @Param author body dto.AuthorRequest true "Данные автора"
// @Success 201 "Автор успешно создан"
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/author [post]
func (h *AuthorHandler) CreateAuthor(c echo.Context) error {
	userID := handler.GetUserId(c)
	var req dto.AuthorRequest
	if err := c.Bind(&req); err != nil {
		return ApiError.NewAPIError(http.StatusBadRequest, "invalid data")
	}
	domainAuthor := fromDto.FromDtoAuthor(&req)
	err := h.service.CreateAuthor(domainAuthor, userID)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusCreated)
}

// UpdateAuthor godoc
// @Summary Обновить автора
// @Description Обновление автора. Требуется авторизация.
// @Tags authors
// @Accept json
// @Produce json
// @Security Bearer
// @Param author body dto.AuthorRequest true "Данные автора"
// @Success 201 "Автор успешно обновлен"
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/author [patch]
func (h *AuthorHandler) UpdateAuthor(c echo.Context) error {
	userID := handler.GetUserId(c)
	var req dto.AuthorRequest
	if err := c.Bind(&req); err != nil {
		return ApiError.NewAPIError(http.StatusBadRequest, "invalid data")
	}
	domainAuthor := fromDto.FromDtoAuthor(&req)
	err := h.service.UpdateAuthor(domainAuthor, userID)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusAccepted)
}

func NewAuthorHandler(service author.AuthorService) *AuthorHandler {
	return &AuthorHandler{service: service}
}
