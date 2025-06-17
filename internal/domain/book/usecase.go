package book

import (
	"go_library/internal/domain/models"
)

type BookService interface {
	GetAllBooks() ([]*models.Book, error)
	GetBookByID(id string) (*models.Book, error)
	CreateBook(book *models.Book) (string, error)
	UpdateBook(id string, book *models.Book) (string, error)
	DeleteBook(id string) error
}
