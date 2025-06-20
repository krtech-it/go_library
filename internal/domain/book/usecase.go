package book

import (
	"go_library/internal/domain/models"
)

type BookService interface {
	GetAllBooks(page, pageSize int) ([]*models.Book, int, error)
	GetBookByID(id string) (*models.Book, error)
	CreateBook(book *models.Book, userId string) (string, error)
	UpdateBook(id string, book *models.Book) (string, error)
	DeleteBook(id string) error
}
