package book

import (
	"errors"
	"github.com/google/uuid"
	domainModel "go_library/internal/domain/models"
	bookRepo "go_library/internal/infrastructure/repository/book"
	"go_library/internal/utils/mapper"
	"time"
)

type bookService struct {
	repo bookRepo.BookRepository
}

func (s *bookService) GetAllBooks() ([]*domainModel.Book, error) {
	books, err := s.repo.GetAllBooks()
	if err != nil {
		return nil, err
	}
	// Convert models to schemas
	result := make([]*domainModel.Book, 0)
	for _, book := range books {
		result = append(result, mapper.ToBookDomain(&book))
	}
	return result, nil
}

func (s *bookService) GetBookByID(id string) (*domainModel.Book, error) {
	book, err := s.repo.GetBookByID(id)
	result := mapper.ToBookDomain(&book)
	return result, err
}

func (s *bookService) CreateBook(book *domainModel.Book) (string, error) {
	bookModel := mapper.FromDomainToBookModel(book)
	bookModel.Id = uuid.NewString()
	if err := s.repo.CheckAuthorByID(bookModel.AuthorID); err != nil {
		return bookModel.Id, errors.New("author not exist")
	}
	if err := s.repo.CheckBookName(bookModel.Title); err == nil {
		return bookModel.Id, errors.New("Book already exists")
	}
	err := s.repo.CreateBook(bookModel)
	return bookModel.Id, err
}

func (s *bookService) UpdateBook(id string, book *domainModel.Book) (string, error) {
	bookDB, err := s.repo.GetBookByID(id)
	if err != nil {
		return id, errors.New("book not exist")
	}
	bookModel := mapper.FromDomainToBookModel(book)
	bookModel.UpdatedAt = time.Time{}
	if err := s.repo.CheckAuthorByID(bookModel.AuthorID); err != nil {
		return id, errors.New("author not exist")
	}
	if bookDB.Title != bookModel.Title {
		if err := s.repo.CheckBookName(bookModel.Title); err == nil {
			return id, errors.New("Book already exists")
		}
	}
	err = s.repo.UpdateBook(id, bookModel)
	return id, err
}

func (s *bookService) DeleteBook(id string) error {
	return s.repo.DeleteBook(id)
}

func NewBookService(repo bookRepo.BookRepository) BookService {
	return &bookService{repo: repo}
}
