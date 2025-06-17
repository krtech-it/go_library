package author

import (
	"go_library/internal/infrastructure/db/models"
	"gorm.io/gorm"
)

type AuthorRepository interface {
	GetAllAuthors() ([]models.Author, error)
}

type authorRepository struct {
	db *gorm.DB
}

func (r *authorRepository) GetAllAuthors() ([]models.Author, error) {
	var authors []models.Author
	err := r.db.Find(&authors).Error
	return authors, err
}

func NewAuthorRepository(db *gorm.DB) AuthorRepository {
	return &authorRepository{db: db}
}
