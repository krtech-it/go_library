package fromDb

import (
	domainModel "go_library/internal/domain/models"
	gormModel "go_library/internal/infrastructure/db/models"
)

func FromDbAuthor(author *gormModel.Author) *domainModel.Author {
	var books []domainModel.Book
	for _, book := range author.Books {
		books = append(books, domainModel.Book{
			Id:          book.Id,
			Title:       book.Title,
			Description: book.Description,
			CountPage:   book.CountPage,
			CreatedAt:   book.CreatedAt,
			UpdatedAt:   book.UpdatedAt,
		})
	}
	return &domainModel.Author{
		Id:        author.Id,
		FirstName: author.FirstName,
		LastName:  author.LastName,
		Books:     books,
		CreatedAt: author.CreatedAt,
		UpdatedAt: author.UpdatedAt,
	}
}
