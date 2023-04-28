package utils

import (
	"math"
	"time"
)

type AmountsBreakdown struct {
	PractitionerPaymentAmount float64 `json:"practitioner_payment_amount"`
	ReimbursementAmount       float64 `json:"reimbursement_amount"`
	PlatformFeeAmount         float64 `json:"platform_fee_amount"`
}

type GetAmountsBreakdownForAppointmentCancellationPayload struct {
	AmountPayed                       float64   `json:"amount_payed" validate:"required gte=0"`
	ScheduledStartDateTime            time.Time `json:"scheduled_start_date_time" validate:"required"`
	HoursThreshold                    float64   `json:"hours_threshold" validate:"required gte=0"`
	PlatformFee                       float64   `json:"platform_fee" validate:"required gte=0 lte=1"`
	CancellationAt                    time.Time `json:"cancellation_at" validate:"required"`
	PercentageToRefundBeforeThreshold float64   `json:"percentage_to_refund_before_threshold" validate:"required gte=0 lte=1"`
	PercentageToRefundAfterThreshold  float64   `json:"percentage_to_refund_after_threshold" validate:"required gte=0 lte=1"`
}

type GetPercentageToReimbursePayload struct {
	CancellationAt                    time.Time `json:"cancellation_at" validate:"required"`
	AppointmentScheduledDateTime      time.Time `json:"appointment_scheduled_date_time" validate:"required"`
	HoursThreshold                    float64   `json:"hours_threshold" validate:"required gte=0"`
	PercentageToRefundBeforeThreshold float64   `json:"percentage_to_refund_before_threshold" validate:"required gte=0 lte=1"`
	PercentageToRefundAfterThreshold  float64   `json:"percentage_to_refund_after_threshold" validate:"required gte=0 lte=1"`
}

func GetAmountsBreakdownForAppointmentCancellation(payload GetAmountsBreakdownForAppointmentCancellationPayload) (AmountsBreakdown, error) {
	isBeforeThreshold := CheckIfIsCancelledBeforeThreshold(payload.CancellationAt, payload.ScheduledStartDateTime, payload.HoursThreshold)
	if isBeforeThreshold {
		return breakdownForPercentage(payload.AmountPayed, payload.PercentageToRefundBeforeThreshold/100, payload.PlatformFee), nil
	}

	return breakdownForPercentage(payload.AmountPayed, payload.PercentageToRefundAfterThreshold/100, payload.PlatformFee), nil
}

func CheckIfIsCancelledBeforeThreshold(cancelledAt time.Time, appointmentScheduledStartDateTime time.Time, threshold float64) bool {
	maximumCancellationTime := appointmentScheduledStartDateTime.Add(-time.Duration(threshold) * time.Hour).UTC()
	if cancelledAt.UTC().Before(maximumCancellationTime) {
		return true
	}

	return false
}

func GetPercentageToReimburse(payload GetPercentageToReimbursePayload) float64 {
	if CheckIfIsCancelledBeforeThreshold(payload.CancellationAt, payload.AppointmentScheduledDateTime, payload.HoursThreshold) {
		return payload.PercentageToRefundBeforeThreshold
	}

	return payload.PercentageToRefundAfterThreshold
}

func breakdownForPercentage(amountPayed, percentage, platformFee float64) AmountsBreakdown {
	percentageToReimburse := percentage
	platformFeePercentage := platformFee

	reimbursementAmount := math.Ceil(amountPayed * percentageToReimburse)
	remainingAmount := amountPayed - reimbursementAmount
	paymentToPractitioner := remainingAmount * (1 - platformFeePercentage)
	paymentToPlatform := remainingAmount * platformFeePercentage

	return AmountsBreakdown{
		PractitionerPaymentAmount: paymentToPractitioner,
		ReimbursementAmount:       reimbursementAmount,
		PlatformFeeAmount:         paymentToPlatform,
	}
}
