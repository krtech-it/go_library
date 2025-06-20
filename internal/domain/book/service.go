package book

import (
	"github.com/google/uuid"
	domainModel "go_library/internal/domain/models"
	ApiError "go_library/internal/errors"
	"go_library/internal/infrastructure/db/mapper/fromDb"
	"go_library/internal/infrastructure/db/mapper/toDb"
	bookRepo "go_library/internal/infrastructure/repository/book"
	"net/http"
	"time"
)

type bookService struct {
	repo bookRepo.BookRepository
}

func (s *bookService) GetAllBooks(page, pageSize int) ([]*domainModel.Book, int, error) {
	books, err := s.repo.GetAllBooks(page, pageSize)
	if err != nil {
		return nil, 0, ApiError.NewAPIError(http.StatusInternalServerError, "Could not get books")
	}
	count, _ := s.repo.GetCountBooks()
	// Convert models to schemas
	result := make([]*domainModel.Book, 0)
	for _, book := range books {
		result = append(result, fromDb.FromDbBook(&book))
	}
	return result, count, nil
}

func (s *bookService) GetBookByID(id string) (*domainModel.Book, error) {
	book, err := s.repo.GetBookByID(id)
	if err != nil {
		return nil, ApiError.NewAPIError(http.StatusNotFound, "Could not get book")
	}
	result := fromDb.FromDbBook(&book)
	return result, nil
}

func (s *bookService) CreateBook(book *domainModel.Book, userId string) (string, error) {
	gormUser, err := s.repo.GetUser(userId)
	if err != nil {
		return "", ApiError.NewAPIError(http.StatusInternalServerError, "Could not get user")
	}
	domainUser := fromDb.FromDbUser(gormUser)
	if domainUser.AuthorID == nil {
		return "", ApiError.NewAPIError(http.StatusConflict, "User not have author")
	}
	book.Author = domainModel.Author{Id: *domainUser.AuthorID}
	bookModel := toDb.ToDbBook(book)
	bookModel.Id = uuid.NewString()
	if err := s.repo.CheckAuthorByID(bookModel.AuthorID); err != nil {
		return bookModel.Id, ApiError.NewAPIError(http.StatusBadRequest, "author not exist")
	}
	if err := s.repo.CheckBookName(bookModel.Title); err == nil {
		return bookModel.Id, ApiError.NewAPIError(http.StatusConflict, "Book already exists")
	}
	err = s.repo.CreateBook(bookModel)
	if err != nil {
		return bookModel.Id, ApiError.NewAPIError(http.StatusInternalServerError, "internal server error")
	}
	return bookModel.Id, nil
}

func (s *bookService) UpdateBook(id string, book *domainModel.Book, userId string) (string, error) {
	gormUser, err := s.repo.GetUser(userId)
	if err != nil {
		return "", ApiError.NewAPIError(http.StatusInternalServerError, "Could not get user")
	}
	domainUser := fromDb.FromDbUser(gormUser)
	bookDB, err := s.repo.GetBookByID(id)
	if err != nil {
		return id, ApiError.NewAPIError(http.StatusNotFound, "Book not found")
	}
	if domainUser.AuthorID == nil {
		return "", ApiError.NewAPIError(http.StatusConflict, "User not have author")
	}
	if *domainUser.AuthorID != bookDB.AuthorID {
		return "", ApiError.NewAPIError(http.StatusForbidden, "User not author")
	}
	book.Author.Id = *domainUser.AuthorID
	bookModel := toDb.ToDbBook(book)
	bookModel.UpdatedAt = time.Time{}
	if err := s.repo.CheckAuthorByID(bookModel.AuthorID); err != nil {
		return id, ApiError.NewAPIError(http.StatusBadRequest, "author not exist")
	}
	if bookDB.Title != bookModel.Title {
		if err := s.repo.CheckBookName(bookModel.Title); err == nil {
			return id, ApiError.NewAPIError(http.StatusConflict, "Book already exists")
		}
	}
	err = s.repo.UpdateBook(id, bookModel)
	if err != nil {
		return id, ApiError.NewAPIError(http.StatusInternalServerError, "internal server error")
	}
	return id, nil
}

func (s *bookService) DeleteBook(id string) error {
	err := s.repo.DeleteBook(id)
	if err != nil {
		return ApiError.NewAPIError(http.StatusInternalServerError, "internal server error")
	}
	return nil
}

func NewBookService(repo bookRepo.BookRepository) BookService {
	return &bookService{repo: repo}
}
