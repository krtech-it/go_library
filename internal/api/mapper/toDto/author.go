package toDto

import (
	"go_library/internal/api/dto"
	domainModel "go_library/internal/domain/models"
)

func ToDtoAuthorMini(a *domainModel.Author) *dto.AuthorResponse {
	return &dto.AuthorResponse{
		Id:        a.Id,
		FirstName: a.FirstName,
		LastName:  a.LastName,
	}
}

func ToDtoAuthor(author *domainModel.Author) *dto.AuthorResponse {
	return ToDtoAuthorMini(author)
}

func ToDtoAuthorWithBooks(author *domainModel.Author) *dto.AuthorFullResponse {
	var books []dto.BaseBookResponse
	for _, book := range author.Books {
		books = append(books, *ToDtoBaseBook(&book))
	}
	return &dto.AuthorFullResponse{
		AuthorResponse: *ToDtoAuthor(author),
		Books:          books,
	}
}
