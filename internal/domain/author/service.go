package author

import (
	domainModel "go_library/internal/domain/models"
	ApiError "go_library/internal/errors"
	autorRepo "go_library/internal/infrastructure/repository/author"
	"go_library/internal/utils/mapper"
	"net/http"
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
		result = append(result, mapper.FromGormToDomainAuthor(&author))
	}
	return result, nil
}

func (s *authorService) GetAuthorByID(id string) (*domainModel.Author, error) {
	author, err := s.repo.GetAuthorByID(id)
	if err != nil {
		return nil, ApiError.NewAPIError(http.StatusNotFound, "Could not get author")
	}
	result := mapper.FromGormToDomainAuthor(author)
	return result, nil
}

func NewAuthorService(repo autorRepo.AuthorRepository) AuthorService {
	return &authorService{repo: repo}
}
