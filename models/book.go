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

type Book struct {
	Id          string    `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CountPage   uint64    `json:"count_page"`
	AuthorID    string    `json:"author_id"`
	Author      Author    `gorm:"foreignKey:AuthorID" json:"author"`
	Genres      []*Genre  `gorm:"many2many:book_genres" json:"genres"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Genre struct {
	Id        string
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
