package mapper

import (
	"go_library/internal/api/dto"
	domainModel "go_library/internal/domain/models"
	gormModel "go_library/internal/infrastructure/db/models"
)

func FromGormToDomainAuthor(author *gormModel.Author) *domainModel.Author {
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

func FromDomainToResponseAuthor(author *domainModel.Author) *dto.AuthorResponse {
	return &dto.AuthorResponse{
		Id:        author.Id,
		FirstName: author.FirstName,
		LastName:  author.LastName,
	}
}

func FromDomainToResponseAuthorFull(author *domainModel.Author) *dto.AuthorFullResponse {
	var books []dto.BaseBookResponse
	for _, book := range author.Books {
		books = append(books, dto.BaseBookResponse{
			Id:          book.Id,
			Title:       book.Title,
			Description: book.Description,
			CountPage:   book.CountPage,
		})
	}
	return &dto.AuthorFullResponse{
		AuthorResponse: *FromDomainToResponseAuthor(author),
		Books:          books,
	}
}

func FromRequestToDomainAuthor(author *dto.AuthorRequest) *domainModel.Author {
	return &domainModel.Author{
		FirstName: author.FirstName,
		LastName:  author.LastName,
	}
}

func FromDomainToGormAuthor(author *domainModel.Author) *gormModel.Author {
	return &gormModel.Author{
		Id:        author.Id,
		FirstName: author.FirstName,
		LastName:  author.LastName,
	}
}
