package models

import (
	"time"
)

type Author struct {
	Id        string
	FirstName string
	LastName  string
	Books     []Book
	CreatedAt time.Time
	UpdatedAt time.Time
}
