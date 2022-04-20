package models

type CreatePaymentRequest struct {
	CommerceOrder   string  `json:"commerce_order"`
	Subject         string  `json:"subject"`
	Currency        string  `json:"currency"`
	Amount          float64 `json:"amount"`
	Email           string  `json:"email"`
	PaymentMethod   int     `json:"payment_method"`
	URLConfirmation string  `json:"url_confirmation"`
	URLReturn       string  `json:"url_return"`
	Optional        string  `json:"optional"`
	Timeout         int     `json:"timeout"`
	MerchantID      string  `json:"merchant_id"`
	PaymentCurrency string  `json:"payment_currency"`
}

type CreatePaymentResponse struct {
	URL                string `json:"url"`
	Token              string `json:"token"`
	FlowOrder          string `json:"flow_order"`
	TimeToPayInSeconds int    `json:"time_to_pay_in_seconds"`
}
