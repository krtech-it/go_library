package fromDto

import (
	"go_library/internal/api/dto"
	domainModel "go_library/internal/domain/models"
)

func FromDtoBook(b *dto.BookRequest) *domainModel.Book {
	return &domainModel.Book{
		Title:       b.Title,
		Description: b.Description,
		CountPage:   b.CountPage,
	}
}
