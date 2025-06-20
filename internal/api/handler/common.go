package handler

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func GetClaims(c echo.Context) jwt.MapClaims {
	user := c.Get("user").(*jwt.Token)
	return user.Claims.(jwt.MapClaims)
}

func GetUserId(c echo.Context) string {
	claims := GetClaims(c)
	return claims["user_id"].(string)
}
