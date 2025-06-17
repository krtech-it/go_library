package book

import (
	"go_library/internal/domain/models"
)

type BookService interface {
	GetAllBooks() ([]*models.Book, error)
	GetBookByID(id string) (*models.Book, error)
	//CreateBook(book schemas.BookRequest) (schemas.BookIdResponse, error)
	//UpdateBook(id string, book schemas.BookRequest) (schemas.BookIdResponse, error)
	//DeleteBook(id string) error
}
