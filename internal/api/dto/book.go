package dto

import "time"

type BaseBookResponse struct {
	Id          string `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
	Title       string `json:"title" example:"The Great Gatsby"`
	Description string `json:"description" example:"A story of the fabulously wealthy Jay Gatsby"`
	CountPage   uint64 `json:"count_page" example:"180"`
}

// BookResponse represents a book response object
// @Description Book response object with detailed information including author and metadata
type BookResponse struct {
	BaseBookResponse `json:",inline"`
	Author           AuthorResponse `json:"author"`
	CreatedAt        time.Time      `json:"created_at" example:"2024-03-20T10:00:00Z"`
	UpdatedAt        time.Time      `json:"updated_at" example:"2024-03-20T10:00:00Z"`
}

// BookRequest represents a book request object
// @Description Book request object for creating or updating a book
type BookRequest struct {
	Title       string `json:"title" binding:"required" example:"The Great Gatsby"`
	Description string `json:"description" binding:"required" example:"A story of the fabulously wealthy Jay Gatsby"`
	CountPage   uint64 `json:"count_page" binding:"required" example:"180"`
	AuthorID    string `json:"author_id" binding:"required" example:"123e4567-e89b-12d3-a456-426614174000"`
}

// BookIdResponse represents a book ID response object
// @Description Book ID response object returned after creation or update
type BookIdResponse struct {
	ID string `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
}
