package author

import (
	"go_library/internal/infrastructure/db/models"
	"gorm.io/gorm"
)

type AuthorRepository interface {
	GetAllAuthors() ([]models.Author, error)
	GetAuthorByID(id string) (*models.Author, error)
	GetUser(id string) (*models.User, error)
	CreateAuthor(author *models.Author) error
	JoinAuthorUser(authorId, userId string) error
	UpdateAuthor(author *models.Author) error
}

type authorRepository struct {
	db *gorm.DB
}

func (r *authorRepository) GetAllAuthors() ([]models.Author, error) {
	var authors []models.Author
	err := r.db.Find(&authors).Error
	return authors, err
}

func (r *authorRepository) GetAuthorByID(id string) (*models.Author, error) {
	var author models.Author
	err := r.db.Preload("Books").Where(map[string]string{"id": id}).First(&author).Error
	return &author, err
}

func (r *authorRepository) GetUser(id string) (*models.User, error) {
	var user models.User
	err := r.db.Preload("Author").Where(map[string]string{"id": id}).First(&user).Error
	return &user, err
}

func (r *authorRepository) CreateAuthor(author *models.Author) error {
	return r.db.Create(&author).Error
}

func (r *authorRepository) JoinAuthorUser(userId string, authorId string) error {
	return r.db.Model(models.User{}).Where(map[string]string{"id": userId}).Update("author_id", authorId).Error
}

func (r *authorRepository) UpdateAuthor(author *models.Author) error {
	return r.db.Model(models.Author{}).Where(map[string]string{"id": author.Id}).Updates(author).Error
}

func NewAuthorRepository(db *gorm.DB) AuthorRepository {
	return &authorRepository{db: db}
}
