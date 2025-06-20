package author

import (
	"github.com/google/uuid"
	domainModel "go_library/internal/domain/models"
	ApiError "go_library/internal/errors"
	"go_library/internal/infrastructure/db/mapper/fromDb"
	"go_library/internal/infrastructure/db/mapper/toDb"
	autorRepo "go_library/internal/infrastructure/repository/author"
	"net/http"
	"time"
)

type authorService struct {
	repo autorRepo.AuthorRepository
}

func (s *authorService) GetAllAuthors() ([]*domainModel.Author, error) {
	authors, err := s.repo.GetAllAuthors()
	if err != nil {
		return nil, ApiError.NewAPIError(http.StatusInternalServerError, "Could not get authors")
	}
	result := make([]*domainModel.Author, 0)
	for _, author := range authors {
		result = append(result, fromDb.FromDbAuthor(&author))
	}
	return result, nil
}

func (s *authorService) GetAuthorByID(id string) (*domainModel.Author, error) {
	author, err := s.repo.GetAuthorByID(id)
	if err != nil {
		return nil, ApiError.NewAPIError(http.StatusNotFound, "Could not get author")
	}
	result := fromDb.FromDbAuthor(author)
	return result, nil
}

func (s *authorService) CreateAuthor(author *domainModel.Author, userId string) error {
	gormUser, err := s.repo.GetUser(userId)
	if err != nil {
		return ApiError.NewAPIError(http.StatusInternalServerError, "Could not get user")
	}
	domainUser := fromDb.FromDbUser(gormUser)
	if domainUser.AuthorID != nil {
		return ApiError.NewAPIError(http.StatusForbidden, "Пользователь уже создал автора")
	}
	gormAuthor := toDb.ToDbAuthor(author)
	gormAuthor.Id = uuid.NewString()
	err = s.repo.CreateAuthor(gormAuthor)
	if err != nil {
		return ApiError.NewAPIError(http.StatusInternalServerError, "Could not create author")
	}
	err = s.repo.JoinAuthorUser(userId, gormAuthor.Id)
	if err != nil {
		return ApiError.NewAPIError(http.StatusInternalServerError, "Could not join author")
	}
	return nil
}

func (s *authorService) UpdateAuthor(author *domainModel.Author, userId string) error {
	gormUser, err := s.repo.GetUser(userId)
	if err != nil {
		return ApiError.NewAPIError(http.StatusInternalServerError, "Could not get user")
	}
	domainUser := fromDb.FromDbUser(gormUser)
	if domainUser.AuthorID == nil {
		return ApiError.NewAPIError(http.StatusForbidden, "У пользователя еще нет автора")
	}
	gormAuthor := toDb.ToDbAuthor(author)
	gormAuthor.Id = *domainUser.AuthorID
	gormAuthor.UpdatedAt = time.Now()
	return s.repo.UpdateAuthor(gormAuthor)
}

func NewAuthorService(repo autorRepo.AuthorRepository) AuthorService {
	return &authorService{repo: repo}
}
