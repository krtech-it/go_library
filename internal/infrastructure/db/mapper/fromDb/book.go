package fromDb

import (
	domainModel "go_library/internal/domain/models"
	gormModel "go_library/internal/infrastructure/db/models"
)

func FromDbBook(b *gormModel.Book) *domainModel.Book {
	var genres []domainModel.Genre
	for _, genre := range b.Genres {
		genres = append(genres, domainModel.Genre{
			Id:   genre.Id,
			Name: genre.Name,
		})
	}
	return &domainModel.Book{
		Id:          b.Id,
		Title:       b.Title,
		Description: b.Description,
		CountPage:   b.CountPage,
		Genres:      genres,
		Author: domainModel.Author{
			Id:        b.Author.Id,
			FirstName: b.Author.FirstName,
			LastName:  b.Author.LastName,
		},
		CreatedAt: b.CreatedAt,
		UpdatedAt: b.UpdatedAt,
	}
}
