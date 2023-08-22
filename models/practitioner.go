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
	Slug                    string     `json:"slug"`
	Rating                  float64    `json:"rating"`
}

type PractitionerAggregated struct {
	ID                  uint64      `json:"id"`
	UID                 string      `json:"uid"`
	FirstName           string      `json:"first_name"`
	SecondName          string      `json:"second_name"`
	LastName            string      `json:"last_name"`
	SecondLastName      string      `json:"second_last_name"`
	FullName            string      `json:"full_name"`
	NationalID          string      `json:"national_id"`
	RegistryID          string      `json:"registry_id"`
	Image               string      `json:"image"`
	Gender              string      `json:"gender"`
	Bio                 string      `json:"bio"`
	PracticeUniversity  *University `json:"practice_university"`
	SpecialtyUniversity *University `json:"specialty_university"`
	Practice            *Practice   `json:"practice"`
	Specialty           *Specialty  `json:"specialty"`
	InternalID          string      `json:"internal_id"`
	Slug                string      `json:"slug"`
	Status              string      `json:"status"`
	Rating              float64     `json:"rating"`
}

type PublicPractitioner struct {
	ID             uint64  `json:"id"`
	UID            string  `json:"uid"`
	FirstName      string  `json:"first_name"`
	SecondName     string  `json:"second_name"`
	LastName       string  `json:"last_name"`
	SecondLastName string  `json:"second_last_name"`
	FullName       string  `json:"full_name"`
	RegistryID     string  `json:"registry_id"`
	Image          string  `json:"image"`
	Gender         string  `json:"gender"`
	Bio            string  `json:"bio"`
	Rating         float64 `json:"rating"`
}

type PublicPractitionerAggregated struct {
	ID                  uint64      `json:"id"`
	UID                 string      `json:"uid"`
	FirstName           string      `json:"first_name"`
	SecondName          string      `json:"second_name"`
	LastName            string      `json:"last_name"`
	SecondLastName      string      `json:"second_last_name"`
	FullName            string      `json:"full_name"`
	RegistryID          string      `json:"registry_id"`
	Image               string      `json:"image"`
	Gender              string      `json:"gender"`
	Bio                 string      `json:"bio"`
	PracticeUniversity  *University `json:"practice_university"`
	SpecialtyUniversity *University `json:"specialty_university"`
	Practice            *Practice   `json:"practice"`
	Specialty           *Specialty  `json:"specialty"`
	Slug                string      `json:"slug"`
	Rating              float64     `json:"rating"`
}

func (p *PublicPractitioner) ToPublicPractitioner() PublicPractitioner {
	return PublicPractitioner{
		ID:             p.ID,
		UID:            p.UID,
		FirstName:      p.FirstName,
		SecondName:     p.SecondName,
		LastName:       p.LastName,
		SecondLastName: p.SecondLastName,
		FullName:       p.FullName,
		RegistryID:     p.RegistryID,
		Image:          p.Image,
		Gender:         p.Gender,
		Bio:            p.Bio,
		Rating:         p.Rating,
	}
}

func (p *PractitionerAggregated) ToPublicPractitionerAggregated() PublicPractitionerAggregated {
	return PublicPractitionerAggregated{
		ID:                  p.ID,
		UID:                 p.UID,
		FirstName:           p.FirstName,
		SecondName:          p.SecondName,
		LastName:            p.LastName,
		SecondLastName:      p.SecondLastName,
		FullName:            p.FullName,
		RegistryID:          p.RegistryID,
		Image:               p.Image,
		Gender:              p.Gender,
		Bio:                 p.Bio,
		PracticeUniversity:  p.PracticeUniversity,
		SpecialtyUniversity: p.SpecialtyUniversity,
		Practice:            p.Practice,
		Specialty:           p.Specialty,
		Slug:                p.Slug,
		Rating:              p.Rating,
	}
}
