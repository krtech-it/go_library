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

func ToBookDomain(b *gormModel.Book) *domainModel.Book {
	return &domainModel.Book{
		Id:          b.Id,
		Title:       b.Title,
		Description: b.Description,
		CountPage:   b.CountPage,
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
		Author: domainModel.Author{
			Id: b.AuthorID,
		},
	}
}
