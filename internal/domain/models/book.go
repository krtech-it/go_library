package models

import "time"

type Book struct {
	Id          string
	Title       string
	Description string
	CountPage   uint64
	Author      Author
	Genres      []Genre
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
