package author

import domainModel "go_library/internal/domain/models"

type AuthorService interface {
	GetAllAuthors() ([]*domainModel.Author, error)
	GetAuthorByID(id string) (*domainModel.Author, error)
	CreateAuthor(author *domainModel.Author, userId string) error
}
