package paging

// PagedResponse is a generic response type wrapping paging logic of the anexia engine.
type PagedResponse[T any] struct {
	Page       int `json:"page"`
	TotalItems int `json:"total_items"`
	TotalPages int `json:"total_pages"`
	Limit      int `json:"limit"`
	Data       []T `json:"data"`
}
