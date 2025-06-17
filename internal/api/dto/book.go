package dto

import "time"

type BookResponse struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CountPage   uint64 `json:"count_page"`
	Author      AuthorResponse
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type BookRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	CountPage   uint64 `json:"count_page"`
	AuthorID    string `json:"author_id"`
}

type BookIdResponse struct {
	ID string `json:"id"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
