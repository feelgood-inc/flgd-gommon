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

func (s *appointmentsTestSuite) TestCheckIfIsCancelledBeforeThreshold_WhenAppointmentIsCancelledBeforeThreshold() {
	cancelledAt := time.Now().UTC()
	startTime := time.Now().UTC().Add(25 * time.Hour)
	threshold := 24.0

	isBeforeThreshold := CheckIfIsCancelledBeforeThreshold(cancelledAt, startTime, threshold)

	s.True(isBeforeThreshold)
}

func (s *appointmentsTestSuite) TestCheckIfIsCancelledBeforeThreshold_WhenAppointmentIsCancelledAfterThreshold() {
	cancelledAt := time.Now().UTC()
	startTime := time.Now().UTC().Add(12 * time.Hour)
	threshold := 24.0

	isBeforeThreshold := CheckIfIsCancelledBeforeThreshold(cancelledAt, startTime, threshold)

	s.False(isBeforeThreshold)
}

func (s *appointmentsTestSuite) TestGetAmountsBreakdownForAppointmentCancellation_WhenAppointmentIsCancelledBeforeThreshold() {
	amountPayed := 30000.0
	scheduledStartDateTime := time.Now().UTC().Add(25 * time.Hour)
	hoursThreshold := 24.0
	platformFee := 0.15
	cancellationAt := time.Now().UTC()

	amountsBreakdown, _ := GetAmountsBreakdownForAppointmentCancellation(GetAmountsBreakdownForAppointmentCancellationPayload{
		AmountPayed:            amountPayed,
		ScheduledStartDateTime: scheduledStartDateTime,
		HoursThreshold:         hoursThreshold,
		PlatformFee:            platformFee,
		CancellationAt:         cancellationAt,
	})

	s.Equal(float64(5100), amountsBreakdown.PractitionerPaymentAmount)
	s.Equal(float64(24000), amountsBreakdown.ReimbursementAmount)
	s.Equal(float64(900), amountsBreakdown.PlatformFeeAmount)
}

func (s *appointmentsTestSuite) TestGetAmountsBreakdownForAppointmentCancellation_WhenAppointmentIsCancelledAfterThreshold() {
	amountPayed := 100.0
	scheduledStartDateTime := time.Now().UTC().Add(12 * time.Hour)
	hoursThreshold := 24.0
	platformFee := 0.1
	cancellationAt := time.Now().UTC()

	amountsBreakdown, _ := GetAmountsBreakdownForAppointmentCancellation(GetAmountsBreakdownForAppointmentCancellationPayload{
		AmountPayed:            amountPayed,
		ScheduledStartDateTime: scheduledStartDateTime,
		HoursThreshold:         hoursThreshold,
		PlatformFee:            platformFee,
		CancellationAt:         cancellationAt,
	})

	s.Equal(float64(81), amountsBreakdown.PractitionerPaymentAmount)
	s.Equal(float64(10), amountsBreakdown.ReimbursementAmount)
	s.Equal(float64(9), amountsBreakdown.PlatformFeeAmount)
}
