package models

type Pagination struct {
	Page         int  `json:"page"`
	ItemsPerPage int  `json:"items_per_page"`
	TotalItems   *int `json:"total_items"`
	TotalPages   *int `json:"total_pages"`
}
