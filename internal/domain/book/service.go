package book

import (
	domainModel "go_library/internal/domain/models"
	bookRepo "go_library/internal/infrastructure/repository/book"
	"go_library/internal/utils/mapper"
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

func NewBookService(repo bookRepo.BookRepository) BookService {
	return &bookService{repo: repo}
}
