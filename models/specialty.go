package models

type Specialty struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	PracticeID int64  `json:"practice_id"`
	Status     string `json:"status"`
}

type SubSpecialty struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	PracticeID  int64  `json:"practice_id"`
	SpecialtyID int64  `json:"specialty_id"`
}
