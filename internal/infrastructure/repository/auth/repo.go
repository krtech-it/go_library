package auth

import (
	"go_library/internal/infrastructure/db/models"
	"gorm.io/gorm"
)

type AuthRepository interface {
	GetUser(username string) (*models.User, error)
}

type authRepository struct {
	db *gorm.DB
}

func (r *authRepository) GetUser(username string) (*models.User, error) {
	var user models.User
	err := r.db.Where(map[string]string{"username": username}).First(&user).Error
	return &user, err
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}
