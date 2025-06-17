package models

import (
	"time"
)

type Author struct {
	Id        string    `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Books     []Book    `gorm:"foreignKey:AuthorID" json:"books"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
