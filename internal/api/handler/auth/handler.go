package auth

import (
	"github.com/labstack/echo/v4"
	"go_library/internal/api/dto"
	"go_library/internal/domain/auth"
	ApiError "go_library/internal/errors"
	"net/http"
)

type AuthHandler struct {
	service auth.AuthService
}

// Login godoc
// @Summary Аутентификация пользователя
// @Description Выполняет вход пользователя и возвращает JWT токен
// @Tags auth
// @Accept x-www-form-urlencoded
// @Produce json
// @Param username formData string true "Имя пользователя" example("admin")
// @Param password formData string true "Пароль" example("1234")
// @Success 200 {object} dto.TokenResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /login [post]
func (h *AuthHandler) Login(c echo.Context) error {
	// Пример проверки логина
	var req dto.AuthLogin
	if err := c.Bind(&req); err != nil {
		return ApiError.NewAPIError(http.StatusBadRequest, "invalid data")
	}
	token, err := h.service.Login(req.Username, req.Password)
	if err != nil {
		return err
	}
	result := dto.TokenResponse{AccessToken: token}
	return c.JSON(http.StatusOK, result)
}

// Register godoc
// @Summary Регистрация пользователя
// @Description Выполняет регистрацию пользователя и возвращает JWT токен
// @Tags auth
// @Accept x-www-form-urlencoded
// @Produce json
// @Param username formData string true "Имя пользователя" example("admin")
// @Param password formData string true "Пароль" example("1234")
// @Success 200 {object} dto.TokenResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /register [post]
func (h *AuthHandler) Register(c echo.Context) error {
	var req dto.AuthLogin
	if err := c.Bind(&req); err != nil {
		return ApiError.NewAPIError(http.StatusBadRequest, "invalid data")
	}
	token, err := h.service.Register(req.Username, req.Password)
	if err != nil {
		return err
	}
	result := dto.TokenResponse{AccessToken: token}
	return c.JSON(http.StatusCreated, result)
}

func NewAuthHandler(service auth.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}
