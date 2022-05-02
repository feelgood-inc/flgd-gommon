package models

import "time"

type Practitioner struct {
	ID                      uint64     `json:"id"`
	UID                     string     `json:"uid"`
	FirstName               string     `json:"first_name"`
	SecondName              string     `json:"second_name"`
	LastName                string     `json:"last_name"`
	SecondLastName          string     `json:"second_last_name"`
	FullName                string     `json:"full_name"`
	Gender                  string     `json:"gender"`
	NationalID              string     `json:"national_id"`
	RegistryID              string     `json:"registry_id"`
	Bio                     string     `json:"bio"`
	PracticeUniversityID    uint64     `json:"practice_university_id"`
	SpecialtyUniversityID   uint64     `json:"specialty_university_id"`
	Image                   string     `json:"image"`
	Email                   string     `json:"email"`
	PracticeID              int64      `json:"practice_id"`
	SpecialtyID             int64      `json:"specialty_id"`
	CreatedAt               *time.Time `json:"created_at"`
	UpdatedAt               *time.Time `json:"updated_at"`
	DeletedAt               *time.Time `json:"deleted_at"`
	InternalID              string     `json:"internal_id"`
	PracticeName            string     `json:"practice_name"`
	SpecialtyName           string     `json:"specialty_name"`
	PracticeUniversityName  string     `json:"practice_university_name"`
	SpecialtyUniversityName string     `json:"specialty_university_name"`
	Status                  string     `json:"status"`
}
