package toDto

import (
	"go_library/internal/api/dto"
	domainModel "go_library/internal/domain/models"
)

func ToDtoBaseBook(b *domainModel.Book) *dto.BaseBookResponse {
	return &dto.BaseBookResponse{
		Id:          b.Id,
		Title:       b.Title,
		Description: b.Description,
		CountPage:   b.CountPage,
	}
}

func ToDtoBook(b *domainModel.Book) *dto.BookResponse {
	return &dto.BookResponse{
		BaseBookResponse: *ToDtoBaseBook(b),
		Author:           *ToDtoAuthorMini(&b.Author),
		CreatedAt:        b.CreatedAt,
		UpdatedAt:        b.UpdatedAt,
	}
}

func ToDtoBooksPaginate(books []*dto.BookResponse, page, pageSize, count int) *dto.BookResponsePagination[dto.BookResponse] {
	return &dto.BookResponsePagination[dto.BookResponse]{
		PageStruct: dto.PageStruct{
			Page:     page,
			PageSize: pageSize,
			Count:    count,
		},
		Items: books,
	}
}

func ToDtoBookWithGenres(b *domainModel.Book) *dto.BookResponseGenres {
	genres := make([]dto.GenreResponse, 0)
	for _, g := range b.Genres {
		genres = append(genres, *ToDtoGanreMini(&g))
	}
	return &dto.BookResponseGenres{
		BookResponse: *ToDtoBook(b),
		Genres:       genres,
	}
}

func ToDtoBookId(bookId string) *dto.BookIdResponse {
	return &dto.BookIdResponse{
		ID: bookId,
	}
}
