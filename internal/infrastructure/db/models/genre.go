package models

import "time"

type Genre struct {
	Id        string    `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
