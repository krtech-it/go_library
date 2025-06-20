package fromDto

import (
	"go_library/internal/api/dto"
	domainModel "go_library/internal/domain/models"
)

func FromDtoAuthor(author *dto.AuthorRequest) *domainModel.Author {
	return &domainModel.Author{
		FirstName: author.FirstName,
		LastName:  author.LastName,
	}
}
