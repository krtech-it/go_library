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
