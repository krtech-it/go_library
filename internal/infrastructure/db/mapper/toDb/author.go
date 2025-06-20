package toDb

import (
	domainModel "go_library/internal/domain/models"
	gormModel "go_library/internal/infrastructure/db/models"
)

func ToDbAuthor(author *domainModel.Author) *gormModel.Author {
	return &gormModel.Author{
		Id:        author.Id,
		FirstName: author.FirstName,
		LastName:  author.LastName,
	}
}
