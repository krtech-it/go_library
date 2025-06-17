package author

import (
	domainModel "go_library/internal/domain/models"
	autorRepo "go_library/internal/infrastructure/repository/author"
	"go_library/internal/utils/mapper"
)

type authorService struct {
	repo autorRepo.AuthorRepository
}

func (a *authorService) GetAllAuthors() ([]*domainModel.Author, error) {
	authors, err := a.repo.GetAllAuthors()
	if err != nil {
		return nil, err
	}
	result := make([]*domainModel.Author, 0)
	for _, author := range authors {
		result = append(result, mapper.FromGormToDomainAuthor(&author))
	}
	return result, nil
}

func NewAuthorService(repo autorRepo.AuthorRepository) AuthorService {
	return &authorService{repo: repo}
}
