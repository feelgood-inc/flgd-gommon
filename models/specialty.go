package models

type Specialty struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	PracticeID int64  `json:"practice_id"`
}
