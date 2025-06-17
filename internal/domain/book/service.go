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
		//result[i] = schemas.BookResponse{
		//	ID:          book.Id,
		//	Title:       book.Title,
		//	Description: book.Description,
		//	CountPage:   book.CountPage,
		//	Author: schemas.AuthorResponse{
		//		ID:        book.Author.Id,
		//		FirstName: book.Author.FirstName,
		//		LastName:  book.Author.LastName,
		//	},
		//	CreatedAt: book.CreatedAt,
		//	UpdatedAt: book.UpdatedAt,
		//}
	}
	return result, nil
}

func NewBookService(repo bookRepo.BookRepository) BookService {
	return &bookService{repo: repo}
}
