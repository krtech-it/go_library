package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type AuthHandler struct {
	service int
}

func (h *AuthHandler) Login(c echo.Context) error {
	// Пример проверки логина
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username != "admin" || password != "1234" {
		return echo.ErrUnauthorized
	}

	// Генерация токена
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = username
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

func NewAuthHandler(service int) *AuthHandler {
	return &AuthHandler{service: service}
}
