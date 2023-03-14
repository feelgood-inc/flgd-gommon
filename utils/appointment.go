package utils

import (
	"time"
)

type AmountsBreakdown struct {
	PractitionerPaymentAmount float64 `json:"practitioner_payment_amount"`
	ReimbursementAmount       float64 `json:"reimbursement_amount"`
	PlatformFeeAmount         float64 `json:"platform_fee_amount"`
}

type GetAmountsBreakdownForAppointmentCancellationPayload struct {
	AmountPayed            float64   `json:"amount_payed" validate:"required gte=0"`
	ScheduledStartDateTime time.Time `json:"scheduled_start_date_time" validate:"required"`
	HoursThreshold         float64   `json:"hours_threshold" validate:"required gte=0"`
	PlatformFee            float64   `json:"platform_fee" validate:"required gte=0 lte=1"`
}

const (
	percentageToReimburseAfterThreshold  = 0.1
	percentageToReimburseBeforeThreshold = 0.8
)

func GetAmountsBreakdownForAppointmentCancellation(payload GetAmountsBreakdownForAppointmentCancellationPayload) (AmountsBreakdown, error) {
	hoursToAppointment := payload.ScheduledStartDateTime.UTC().Sub(time.Now().UTC()).Hours()
	if hoursToAppointment < payload.HoursThreshold {
		return breakdownForPercentage(payload.AmountPayed, percentageToReimburseAfterThreshold, payload.PlatformFee), nil
	}

	return breakdownForPercentage(payload.AmountPayed, percentageToReimburseBeforeThreshold, payload.PlatformFee), nil
}

func CheckIfIsCancelledBeforeThreshold(appointmentScheduledStartDateTime time.Time, threshold float64) bool {
	hoursToAppointment := appointmentScheduledStartDateTime.UTC().Sub(time.Now().UTC()).Hours()
	if hoursToAppointment < threshold {
		return true
	}

	return false
}

func breakdownForPercentage(amountPayed, percentage, platformFee float64) AmountsBreakdown {
	percentageToReimburse := percentage
	platformFeePercentage := platformFee

	reimbursementAmount := amountPayed * percentageToReimburse
	remainingAmount := amountPayed - reimbursementAmount
	paymentToPractitioner := remainingAmount * (1 - platformFeePercentage)
	paymentToPlatform := remainingAmount * platformFeePercentage

	return AmountsBreakdown{
		PractitionerPaymentAmount: paymentToPractitioner,
		ReimbursementAmount:       reimbursementAmount,
		PlatformFeeAmount:         paymentToPlatform,
	}
}
