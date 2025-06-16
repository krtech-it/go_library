package schemas

import "time"

type BookResponse struct {
	ID          string         `json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	CountPage   uint64         `json:"count_page"`
	Author      AuthorResponse `json:"author"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

type AuthorResponse struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type BookIdResponse struct {
	ID string `json:"id"`
}

type BookRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	CountPage   uint64 `json:"count_page"`
	AuthorID    string `json:"author_id"`
}
