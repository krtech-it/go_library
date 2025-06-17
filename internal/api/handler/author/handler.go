package author

import (
	"github.com/labstack/echo/v4"
	"go_library/internal/api/dto"
	"go_library/internal/domain/author"
	"go_library/internal/utils/mapper"
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
// @Success 200 {array} dto.AuthorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/author [get]
func (h *AuthorHandler) GetAllAuthors(c echo.Context) error {
	authors, err := h.service.GetAllAuthors()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Could not get authors"})
	}
	response := make([]*dto.AuthorResponse, 0)
	for _, value := range authors {
		response = append(response, mapper.FromDomainToResponseAuthor(value))
	}
	return c.JSON(http.StatusOK, response)
}

// GetAuthorByID godoc
// @Summary Получить автора по ID
// @Description Возвращает информацию об авторе по его идентификатору, включая список его книг
// @Tags authors
// @Accept json
// @Produce json
// @Param id path string true "Author ID" example("123e4567-e89b-12d3-a456-426614174000")
// @Success 200 {object} dto.AuthorFullResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/author/{id} [get]
func (h *AuthorHandler) GetAuthorByID(c echo.Context) error {
	id := c.Param("id")
	author, err := h.service.GetAuthorByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Could not get author"})
	}
	response := mapper.FromDomainToResponseAuthorFull(author)
	return c.JSON(http.StatusOK, response)
}

func NewAuthorHandler(service author.AuthorService) *AuthorHandler {
	return &AuthorHandler{service: service}
}
