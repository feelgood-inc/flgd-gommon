package models

import "time"

type Patient struct {
	ID             int64      `json:"id"`
	UID            string     `json:"uid"`
	FirstName      string     `json:"first_name"`
	SecondName     string     `json:"second_name"`
	LastName       string     `json:"last_name"`
	SecondLastName string     `json:"second_last_name"`
	NationalID     string     `json:"national_id"`
	Email          string     `json:"email"`
	Cellphone      string     `json:"cellphone"`
	FullName       string     `json:"full_name"`
	Gender         string     `json:"gender"`
	CreatedAt      *time.Time `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at"`
}
