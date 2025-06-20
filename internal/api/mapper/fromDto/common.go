package fromDto

import domainModel "go_library/internal/domain/models"

func FromDtoPagination(page, pageSize int) *domainModel.Pagination {
	return &domainModel.Pagination{
		Page:     page,
		PageSize: pageSize,
	}
}
