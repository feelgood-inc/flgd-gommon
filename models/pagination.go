package models

type Pagination struct {
	Page         int `json:"page"`
	ItemsPerPage int `json:"items_per_page"`
}
