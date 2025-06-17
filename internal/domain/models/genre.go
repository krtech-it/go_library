package models

import "time"

type Genre struct {
	Id        string
	Name      string
	Books     []Book
	CreatedAt time.Time
	UpdatedAt time.Time
}
