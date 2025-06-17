package dto

// AuthorResponse represents an author response object
// @Description Author response object with basic information
type AuthorResponse struct {
	Id        string `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
	FirstName string `json:"first_name" example:"John"`
	LastName  string `json:"last_name" example:"Doe"`
}

type AuthorFullResponse struct {
	AuthorResponse `json:",inline"`
	Books          []BaseBookResponse `json:"books"`
}
