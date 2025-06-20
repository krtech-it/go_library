package toDb

import (
	domainModel "go_library/internal/domain/models"
	gormModel "go_library/internal/infrastructure/db/models"
)

func ToDbBook(b *domainModel.Book) *gormModel.Book {
	return &gormModel.Book{
		Title:       b.Title,
		Description: b.Description,
		CountPage:   b.CountPage,
		AuthorID:    b.Author.Id,
		CreatedAt:   b.CreatedAt,
		UpdatedAt:   b.UpdatedAt,
	}
}
