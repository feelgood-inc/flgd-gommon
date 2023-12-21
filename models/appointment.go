package models

import "time"

type Appointment struct {
	// Group the int64 fields together
	ID                           int64 `json:"id"`
	DurationInSeconds            int64 `json:"duration_in_seconds"`
	ElapsedTimeInSeconds         int64 `json:"elapsed_time_in_seconds"`
	CancellationThresholdInHours int64 `json:"cancellation_threshold_in_hours"`
	HealthCareFacilityID         int64 `json:"health_care_facility_id"`

	// Group the *time.Time fields together
	ScheduledStartDateTime    *time.Time `json:"scheduled_start_date_time"`
	ScheduledEndDateTime      *time.Time `json:"scheduled_end_date_time"`
	EndedAt                   *time.Time `json:"ended_at"`
	CreatedAt                 *time.Time `json:"created_at"`
	UpdatedAt                 *time.Time `json:"updated_at"`
	BookedAt                  *time.Time `json:"booked_at"`
	ReservedAt                *time.Time `json:"reserved_at"`
	CancellationThresholdTime time.Time  `json:"cancellation_threshold_time"`

	// Group the float64 fields together
	Price                                       float64 `json:"price"`
	PlatformFeeAsPercentage                     float64 `json:"platform_fee_as_percentage"`
	PlatformFeeAsFixedAmount                    float64 `json:"platform_fee_as_fixed_amount"`
	RefundPercentageBeforeCancellationThreshold float64 `json:"refund_percentage_before_cancellation_threshold"`
	RefundPercentageAfterCancellationThreshold  float64 `json:"refund_percentage_after_cancellation_threshold"`

	// Group the *int64 fields together
	PatientID                    *int64 `json:"patient_id"`
	PractitionerID               *int64 `json:"practitioner_id"`
	PracticeID                   *int64 `json:"practice_id"`
	RescheduledToAppointmentID   *int64 `json:"rescheduled_to_appointment_id"`
	RescheduledFromAppointmentID *int64 `json:"rescheduled_from_appointment_id"`

	// Group the *string fields together
	PractitionerUID *string `json:"practitioner_uid"`
	PatientUID      *string `json:"patient_uid"`

	// Group the string fields together
	Status     string `json:"status"`
	InternalID string `json:"internal_id"`
	BatchID    string `json:"batch_id"`
	Timezone   string `json:"timezone"`
	Type       string `json:"type"`

	// Group the *struct fields together
	Practitioner *Practitioner `json:"practitioner,omitempty"`
	Patient      *User         `json:"patient,omitempty"`
}

type AppointmentAggregated struct {
	// Similar grouping for AppointmentAggregated
	ID                           int64 `json:"id"`
	DurationInSeconds            int64 `json:"duration_in_seconds"`
	ElapsedTimeInSeconds         int64 `json:"elapsed_time_in_seconds"`
	CancellationThresholdInHours int64 `json:"cancellation_threshold_in_hours"`
	HealthCareFacilityID         int64 `json:"health_care_facility_id"`

	ScheduledStartDateTime    *time.Time `json:"scheduled_start_date_time"`
	ScheduledEndDateTime      *time.Time `json:"scheduled_end_date_time"`
	EndedAt                   *time.Time `json:"ended_at"`
	CreatedAt                 *time.Time `json:"created_at"`
	UpdatedAt                 *time.Time `json:"updated_at"`
	BookedAt                  *time.Time `json:"booked_at"`
	ReservedAt                *time.Time `json:"reserved_at"`
	CancellationThresholdTime time.Time  `json:"cancellation_threshold_time"`

	Price                                       float64 `json:"price"`
	PlatformFeeAsPercentage                     float64 `json:"platform_fee_as_percentage"`
	PlatformFeeAsFixedAmount                    float64 `json:"platform_fee_as_fixed_amount"`
	RefundPercentageBeforeCancellationThreshold float64 `json:"refund_percentage_before_cancellation_threshold"`
	RefundPercentageAfterCancellationThreshold  float64 `json:"refund_percentage_after_cancellation_threshold"`

	Patient            *User               `json:"patient"`
	Practitioner       *Practitioner       `json:"practitioner"`
	HealthCareFacility *HealthCareFacility `json:"health_care_facility"`

	Status          string `json:"status"`
	InternalID      string `json:"internal_id"`
	BatchID         string `json:"batch_id"`
	Timezone        string `json:"timezone"`
	PractitionerUID string `json:"practitioner_uid"`
	PatientUID      string `json:"patient_uid"`
	Type            string `json:"type"`
}
