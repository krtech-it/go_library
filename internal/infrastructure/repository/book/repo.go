package book

import (
	"go_library/internal/infrastructure/db/models"
	"gorm.io/gorm"
)

type BookRepository interface {
	GetAllBooks() ([]models.Book, error)
	GetBookByID(id string) (models.Book, error)
	CreateBook(book *models.Book) error
	CheckAuthorByID(id string) error
	CheckBookName(name string) error
	UpdateBook(id string, book models.Book) error
	DeleteBook(id string) error
}

type bookRepository struct {
	db *gorm.DB
}

func (r *bookRepository) DeleteBook(id string) error {
	return r.db.Where(map[string]interface{}{"id": id}).Delete(&models.Book{}).Error
}

func (r *bookRepository) GetAllBooks() ([]models.Book, error) {
	var books []models.Book
	err := r.db.Preload("Author").Find(&books).Error
	return books, err
}

func (r *bookRepository) GetBookByID(id string) (models.Book, error) {
	var book models.Book
	err := r.db.Preload("Author").Where(map[string]interface{}{"id": id}).Find(&book).Error
	return book, err
}

func (r *bookRepository) CreateBook(book *models.Book) error {
	return r.db.Create(book).Error
}

func (r *bookRepository) CheckAuthorByID(id string) error {
	var author models.Author
	return r.db.Where(map[string]interface{}{"id": id}).First(&author).Error
}

func (r *bookRepository) CheckBookName(name string) error {
	var book models.Book
	return r.db.Where(map[string]interface{}{"title": name}).First(&book).Error
}

func (r *bookRepository) UpdateBook(id string, book models.Book) error {
	return r.db.Model(models.Book{}).Where(map[string]interface{}{"id": id}).Updates(book).Error
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db: db}
}
