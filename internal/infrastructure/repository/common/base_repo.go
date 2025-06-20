package common

import "gorm.io/gorm"

type BaseRepository struct {
	Db *gorm.DB
}

func (r *BaseRepository) Paginate(page, pageSize int) *gorm.DB {
	if page < 1 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 100
	}
	offset := (page - 1) * pageSize
	return r.Db.Offset(offset).Limit(pageSize)
}

func NewBaseRepository(db *gorm.DB) *BaseRepository {
	return &BaseRepository{Db: db}
}
