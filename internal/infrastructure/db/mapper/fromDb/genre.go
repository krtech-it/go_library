package fromDb

import (
	domainModel "go_library/internal/domain/models"
	gormModel "go_library/internal/infrastructure/db/models"
)

func FromDbGanreMini(g *gormModel.Genre) *domainModel.Genre {
	return &domainModel.Genre{
		Id:        g.Id,
		Name:      g.Name,
		CreatedAt: g.CreatedAt,
		UpdatedAt: g.UpdatedAt,
	}
}
