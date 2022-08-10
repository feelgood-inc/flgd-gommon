package models

import "time"

type Appointment struct {
	ID                     int64         `json:"id"`
	PatientID              *int64        `json:"patient_id"`
	PractitionerID         *int64        `json:"practitioner_id"`
	ScheduledStartDateTime *time.Time    `json:"scheduled_start_date_time"`
	ScheduledEndDateTime   *time.Time    `json:"scheduled_end_date_time"`
	Status                 string        `json:"status"`
	EndedAt                *time.Time    `json:"ended_at"`
	AvailableInsurances    *[]string     `json:"available_insurances"`
	AvailableServices      *[]string     `json:"available_services"`
	InternalID             string        `json:"internal_id"`
	DurationInSeconds      int64         `json:"duration_in_seconds"`
	ElapsedTimeInSeconds   int64         `json:"elapsed_time_in_seconds"`
	CreatedAt              *time.Time    `json:"created_at"`
	UpdatedAt              *time.Time    `json:"updated_at"`
	BookedAt               *time.Time    `json:"booked_at"`
	ReservedAt             *time.Time    `json:"reserved_at"`
	Price                  float64       `json:"price"`
	BatchID                string        `json:"batch_id"`
	Timezone               string        `json:"timezone"`
	PractitionerUID        *string       `json:"practitioner_uid"`
	PatientUID             *string       `json:"patient_uid"`
	Practitioner           *Practitioner `json:"practitioner,omitempty"`
	Patient                *User         `json:"patient,omitempty"`
	PracticeID             *int64        `json:"practice_id"`
}

type AppointmentAggregated struct {
	ID                     int64         `json:"id"`
	Patient                *User         `json:"patient"`
	Practitioner           *Practitioner `json:"practitioner"`
	ScheduledStartDateTime *time.Time    `json:"scheduled_start_date_time"`
	ScheduledEndDateTime   *time.Time    `json:"scheduled_end_date_time"`
	Status                 string        `json:"status"`
	EndedAt                *time.Time    `json:"ended_at"`
	AvailableInsurances    *[]string     `json:"available_insurances"`
	AvailableServices      *[]string     `json:"available_services"`
	InternalID             string        `json:"internal_id"`
	DurationInSeconds      int64         `json:"duration_in_seconds"`
	ElapsedTimeInSeconds   int64         `json:"elapsed_time_in_seconds"`
	CreatedAt              *time.Time    `json:"created_at"`
	UpdatedAt              *time.Time    `json:"updated_at"`
	BookedAt               *time.Time    `json:"booked_at"`
	ReservedAt             *time.Time    `json:"reserved_at"`
	Price                  float64       `json:"price"`
	BatchID                string        `json:"batch_id"`
	Timezone               string        `json:"timezone"`
	PractitionerUID        string        `json:"practitioner_uid"`
	PatientUID             string        `json:"patient_uid"`
}
