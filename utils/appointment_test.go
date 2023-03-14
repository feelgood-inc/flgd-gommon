package utils

import (
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type appointmentsTestSuite struct {
	suite.Suite
}

func TestAppointmentsTestSuite(t *testing.T) {
	suite.Run(t, new(appointmentsTestSuite))
}

func (s *appointmentsTestSuite) TestGetAmountsBreakdownForAppointmentCancellation_WhenAppointmentIsCancelledBeforeThreshold() {
	cancelledAt := time.Now().UTC()
	startTime := time.Now().UTC().Add(25 * time.Hour)
	threshold := 24.0

	isBeforeThreshold := CheckIfIsCancelledBeforeThreshold(cancelledAt, startTime, threshold)

	s.True(isBeforeThreshold)
}

func (s *appointmentsTestSuite) TestGetAmountsBreakdownForAppointmentCancellation_WhenAppointmentIsCancelledAfterThreshold() {
	cancelledAt := time.Now().UTC()
	startTime := time.Now().UTC().Add(12 * time.Hour)
	threshold := 24.0

	isBeforeThreshold := CheckIfIsCancelledBeforeThreshold(cancelledAt, startTime, threshold)

	s.False(isBeforeThreshold)
}
