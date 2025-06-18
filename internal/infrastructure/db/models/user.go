package models

import "time"

type User struct {
	Id        string    `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Username  string    `gorm:"unique:true" json:"username"`
	Admin     bool      `json:"admin" gorm:"default:false"`
	Password  string    `json:"password"`
	AuthorID  string    `json:"author_id"`
	Author    Author    `gorm:"foreignKey:AuthorID" json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
