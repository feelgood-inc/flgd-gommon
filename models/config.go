package models

type GlobalConfig struct {
	Thresholds  ThresholdConfig
	Commissions CommissionConfig
}

type ThresholdConfig struct {
	AppointmentCancellationThresholdInHours int
}

type CommissionConfig struct {
	PlatformConsultationCommission float64
}
