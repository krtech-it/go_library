package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"go_library/internal/core"
	ApiError "go_library/internal/errors"
	authRepo "go_library/internal/infrastructure/repository/auth"
	"go_library/internal/utils/mapper"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type authService struct {
	repo authRepo.AuthRepository
}

func (a *authService) Login(username, password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	userGorm, err := a.repo.GetUser(username)
	if err != nil {
		return "", ApiError.NewAPIError(http.StatusNotFound, "Could not get user")
	}
	user := mapper.FromGormToDomainUser(userGorm)
	if user.Password != string(hashPassword) {
		return "", ApiError.NewAPIError(http.StatusUnauthorized, "Invalid username or password")
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = user.Username
	claims["admin"] = user.Admin
	claims["exp"] = time.Now().Add(time.Minute * 40).Unix()
	return token.SignedString(core.JwtSecret)
}

func NewAuthService(repo authRepo.AuthRepository) AuthService {
	return &authService{repo: repo}
}
