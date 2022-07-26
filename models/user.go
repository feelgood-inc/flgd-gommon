package models

type User struct {
	UID            string   `json:"uid"`
	FirstName      string   `json:"first_name"`
	SecondName     string   `json:"second_name"`
	LastName       string   `json:"last_name"`
	SecondLastName string   `json:"second_last_name"`
	Gender         string   `json:"gender"`
	NationalID     string   `json:"national_id"`
	Email          string   `json:"email"`
	Provider       string   `json:"provider"`
	Type           string   `json:"type"`
	Roles          []string `json:"roles"`
}
