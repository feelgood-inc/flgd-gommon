package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAppointmentModel(t *testing.T) {
	appointment := Appointment{
		ID:                     0,
		PatientID:              nil,
		PractitionerID:         nil,
		ScheduledStartDateTime: nil,
		ScheduledEndDateTime:   nil,
		Status:                 "",
		EndedAt:                nil,
		AvailableInsurances:    nil,
		AvailableServices:      nil,
		InternalID:             "",
		DurationInSeconds:      0,
		ElapsedTimeInSeconds:   0,
		CreatedAt:              nil,
		UpdatedAt:              nil,
		BookedAt:               nil,
		ReservedAt:             nil,
		Price:                  32000,
		BatchID:                "",
		Timezone:               "",
		PractitionerUID:        nil,
		PatientUID:             nil,
	}

	assert.IsType(t, Appointment{}, appointment)
}
