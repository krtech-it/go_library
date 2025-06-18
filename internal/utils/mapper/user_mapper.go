package mapper

import (
	domainModel "go_library/internal/domain/models"
	gormModel "go_library/internal/infrastructure/db/models"
)

func FromGormToDomainUser(user *gormModel.User) *domainModel.User {
	return &domainModel.User{
		Id:        user.Id,
		Username:  user.Username,
		Password:  user.Password,
		Admin:     user.Admin,
		FirstName: user.Author.FirstName,
		LastName:  user.Author.LastName,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
