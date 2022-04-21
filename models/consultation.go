package models

import "time"

type Consultation struct {
	ID                   uint64     `json:"id"`
	AppointmentID        uint64     `json:"appointment_id"`
	PatientID            uint64     `json:"patient_id"`
	PractitionerID       uint64     `json:"practitioner_id"`
	Status               string     `json:"status"`
	ElapsedTimeInSeconds int64      `json:"elapsed_time_in_seconds"`
	PaymentStatus        string     `json:"payment_status"`
	CreatedAt            *time.Time `json:"created_at"`
	UpdatedAt            *time.Time `json:"updated_at"`
}
