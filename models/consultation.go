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
	PractitionerUID      string     `json:"practitioner_uid"`
	PatientUID           string     `json:"patient_uid"`
}

type ConsultationAggregated struct {
	ID                   uint64       `json:"id"`
	Appointment          Appointment  `json:"appointment"`
	Patient              Patient      `json:"patient"`
	Practitioner         Practitioner `json:"practitioner"`
	Status               string       `json:"status"`
	ElapsedTimeInSeconds int64        `json:"elapsed_time_in_seconds"`
	PaymentStatus        string       `json:"payment_status"`
	PractitionerUID      string       `json:"practitioner_uid"`
	PatientUID           string       `json:"patient_uid"`
	CreatedAt            *time.Time   `json:"created_at"`
	UpdatedAt            *time.Time   `json:"updated_at"`
}
