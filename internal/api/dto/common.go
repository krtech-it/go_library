package dto

type PageStruct struct {
	Page     int `json:"page" example:"1"`
	PageSize int `json:"page_size" example:"10"`
	Count    int `json:"count" example:"0"`
}

type BookResponsePagination[T any] struct {
	PageStruct `json:",inline"`
	Items      []*T `json:"items"`
}
