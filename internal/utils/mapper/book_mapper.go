package mapper

import (
	"go_library/internal/api/dto"
	domainModel "go_library/internal/domain/models"
	gormModel "go_library/internal/infrastructure/db/models"
)

func FromDomainToResponseBaseBook(b *domainModel.Book) *dto.BaseBookResponse {
	return &dto.BaseBookResponse{
		Id:          b.Id,
		Title:       b.Title,
		Description: b.Description,
		CountPage:   b.CountPage,
	}
}

func ToBookResponse(b *domainModel.Book) *dto.BookResponse {
	return &dto.BookResponse{
		BaseBookResponse: *FromDomainToResponseBaseBook(b),
		Author: dto.AuthorResponse{
			Id:        b.Author.Id,
			FirstName: b.Author.FirstName,
			LastName:  b.Author.LastName,
		},
		CreatedAt: b.CreatedAt,
		UpdatedAt: b.UpdatedAt,
	}
}

func ToBookResponseWithGenres(b *domainModel.Book) *dto.BookResponseGenres {
	genres := make([]dto.GenreResponse, 0)
	for _, g := range b.Genres {
		genres = append(genres, dto.GenreResponse{
			Id:   g.Id,
			Name: g.Name,
		})
	}
	return &dto.BookResponseGenres{
		BookResponse: *ToBookResponse(b),
		Genres:       genres,
	}
}

func ToBookDomain(b *gormModel.Book) *domainModel.Book {
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

func FromDomainToBookModel(b *domainModel.Book) *gormModel.Book {
	return &gormModel.Book{
		Title:       b.Title,
		Description: b.Description,
		CountPage:   b.CountPage,
		AuthorID:    b.Author.Id,
		CreatedAt:   b.CreatedAt,
		UpdatedAt:   b.UpdatedAt,
	}
}

func FromRequestToDomainBook(b *dto.BookRequest) *domainModel.Book {
	return &domainModel.Book{
		Title:       b.Title,
		Description: b.Description,
		CountPage:   b.CountPage,
	}
}
