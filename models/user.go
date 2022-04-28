package models

type User struct {
	UID      string   `json:"uid"`
	Email    string   `json:"email"`
	Provider string   `json:"provider"`
	Type     string   `json:"type"`
	Roles    []string `json:"roles"`
}
