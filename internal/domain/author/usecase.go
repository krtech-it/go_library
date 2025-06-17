package author

import domainModel "go_library/internal/domain/models"

type AuthorService interface {
	GetAllAuthors() ([]*domainModel.Author, error)
}
