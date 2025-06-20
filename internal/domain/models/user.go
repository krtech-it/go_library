package models

import "time"

type User struct {
	Id        string
	Username  string
	FirstName string
	LastName  string
	Admin     bool
	AuthorID  *string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
