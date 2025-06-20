package fromDb

import (
	domainModel "go_library/internal/domain/models"
	gormModel "go_library/internal/infrastructure/db/models"
)

func FromDbUser(u *gormModel.User) *domainModel.User {
	return &domainModel.User{
		Id:        u.Id,
		Username:  u.Username,
		Admin:     u.Admin,
		Password:  u.Password,
		AuthorID:  u.AuthorID,
		FirstName: u.Author.FirstName,
		LastName:  u.Author.LastName,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
