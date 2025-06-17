package author

import (
	domainModel "go_library/internal/domain/models"
	autorRepo "go_library/internal/infrastructure/repository/author"
	"go_library/internal/utils/mapper"
)

type authorService struct {
	repo autorRepo.AuthorRepository
}

func (s *authorService) GetAllAuthors() ([]*domainModel.Author, error) {
	authors, err := s.repo.GetAllAuthors()
	if err != nil {
		return nil, err
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
		return nil, err
	}
	result := mapper.FromGormToDomainAuthor(author)
	return result, err
}

func NewAuthorService(repo autorRepo.AuthorRepository) AuthorService {
	return &authorService{repo: repo}
}
