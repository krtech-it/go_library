package toDto

import (
	"go_library/internal/api/dto"
	domainModel "go_library/internal/domain/models"
)

func ToDtoGanreMini(g *domainModel.Genre) *dto.GenreResponse {
	return &dto.GenreResponse{
		Id:   g.Id,
		Name: g.Name,
	}
}
