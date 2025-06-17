package service

import (
	"errors"
	"github.com/google/uuid"
	"go_library/models"
	"go_library/repository"
	"go_library/schemas"
	"time"
)

//type BookService interface {
//	GetAllBooks() ([]schemas.BookResponse, error)
//	GetBookByID(id string) (schemas.BookResponse, error)
//	CreateBook(book schemas.BookRequest) (schemas.BookIdResponse, error)
//	UpdateBook(id string, book schemas.BookRequest) (schemas.BookIdResponse, error)
//	DeleteBook(id string) error
//}

type bookService struct {
	repo repository.BookRepository
}

func (s *bookService) GetAllBooks() ([]schemas.BookResponse, error) {
	books, err := s.repo.GetAllBooks()
	if err != nil {
		return nil, err
	}
	// Convert models to schemas
	result := make([]schemas.BookResponse, len(books))
	for i, book := range books {
		result[i] = schemas.BookResponse{
			ID:          book.Id,
			Title:       book.Title,
			Description: book.Description,
			CountPage:   book.CountPage,
			Author: schemas.AuthorResponse{
				ID:        book.Author.Id,
				FirstName: book.Author.FirstName,
				LastName:  book.Author.LastName,
			},
			CreatedAt: book.CreatedAt,
			UpdatedAt: book.UpdatedAt,
		}
	}
	return result, nil
}

func (s *bookService) GetBookByID(id string) (schemas.BookResponse, error) {
	book, err := s.repo.GetBookByID(id)
	if err != nil {
		return schemas.BookResponse{}, err
	}
	result := schemas.BookResponse{
		ID:          book.Id,
		Title:       book.Title,
		Description: book.Description,
		CountPage:   book.CountPage,
		Author: schemas.AuthorResponse{
			ID:        book.Author.Id,
			FirstName: book.Author.FirstName,
			LastName:  book.Author.LastName,
		},
		CreatedAt: book.CreatedAt,
		UpdatedAt: book.UpdatedAt,
	}
	return result, nil
}

func (s *bookService) CreateBook(book schemas.BookRequest) (schemas.BookIdResponse, error) {
	bookModel := models.Book{
		Id:          uuid.NewString(),
		Title:       book.Title,
		Description: book.Description,
		CountPage:   book.CountPage,
		AuthorID:    book.AuthorID,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}
	bookId := schemas.BookIdResponse{ID: bookModel.Id}
	if err := s.repo.CheckAuthorByID(book.AuthorID); err != nil {
		return bookId, errors.New("author not exist")
	}
	if err := s.repo.CheckBookName(book.Title); err == nil {
		return bookId, errors.New("Book already exists")
	}
	err := s.repo.CreateBook(bookModel)
	return bookId, err
}

func (s *bookService) UpdateBook(id string, book schemas.BookRequest) (schemas.BookIdResponse, error) {
	bookDB, err := s.repo.GetBookByID(id)
	if err != nil {
		return schemas.BookIdResponse{}, errors.New("book not exist")
	}
	bookModel := models.Book{
		Title:       book.Title,
		Description: book.Description,
		CountPage:   book.CountPage,
		AuthorID:    book.AuthorID,
		UpdatedAt:   time.Time{},
	}
	bookId := schemas.BookIdResponse{ID: id}
	if err := s.repo.CheckAuthorByID(book.AuthorID); err != nil {
		return bookId, errors.New("author not exist")
	}
	if bookDB.Title != bookModel.Title {
		if err := s.repo.CheckBookName(book.Title); err == nil {
			return bookId, errors.New("Book already exists")
		}
	}
	err = s.repo.UpdateBook(id, bookModel)
	return bookId, err
}

func (s *bookService) DeleteBook(id string) error {
	return s.repo.DeleteBook(id)

}

func NewBookService(repo repository.BookRepository) BookService {
	return &bookService{repo: repo}
}
