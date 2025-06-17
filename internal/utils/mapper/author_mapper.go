package mapper

import (
	"go_library/internal/api/dto"
	domainModel "go_library/internal/domain/models"
	gormModel "go_library/internal/infrastructure/db/models"
)

func FromGormToDomainAuthor(author *gormModel.Author) *domainModel.Author {
	return &domainModel.Author{
		Id:        author.Id,
		FirstName: author.FirstName,
		LastName:  author.LastName,
		CreatedAt: author.CreatedAt,
		UpdatedAt: author.UpdatedAt,
	}
}

func FromDomainToResponseAuthor(author *domainModel.Author) *dto.AuthorResponse {
	return &dto.AuthorResponse{
		Id:        author.Id,
		FirstName: author.FirstName,
		LastName:  author.LastName,
	}
}
