package models

import "time"

type CreatePaymentRequest struct {
	CommerceOrder   string              `json:"commerce_order"`
	Subject         string              `json:"subject"`
	Currency        string              `json:"currency"`
	Amount          float64             `json:"amount"`
	Email           string              `json:"email"`
	PaymentMethod   int                 `json:"payment_method"`
	URLConfirmation string              `json:"url_confirmation"`
	URLReturn       string              `json:"url_return"`
	Optional        GatewayOptionalData `json:"optional"`
	Timeout         int                 `json:"timeout"`
	MerchantID      string              `json:"merchant_id"`
	PaymentCurrency string              `json:"payment_currency"`
}

type CreatePaymentResponse struct {
	URL                string `json:"url"`
	Token              string `json:"token"`
	FlowOrder          string `json:"flow_order"`
	TimeToPayInSeconds int    `json:"time_to_pay_in_seconds"`
}

type GatewayOptionalData struct {
	ResourcePayed string      `json:"resource_payed" bson:"resource_payed"`
	ResourceID    string      `json:"resource_id" bson:"resource_id"`
	Data          interface{} `json:"data" bson:"data"`
}

type Payment struct {
	Status                            uint        `json:"status" bson:"status"`
	Amount                            float64     `json:"amount" bson:"amount"`
	Currency                          string      `json:"currency" bson:"currency"`
	PayerEmail                        string      `json:"payer_email" bson:"payer_email"`
	Data                              interface{} `json:"data" bson:"data"`
	PaymentMethod                     string      `json:"payment_method" bson:"payment_method"`
	ConversionDate                    time.Time   `json:"conversion_date" bson:"conversion_date"`
	ConversionRate                    float64     `json:"conversion_rate" bson:"conversion_rate"`
	GatewayFee                        float64     `json:"gateway_fee" bson:"gateway_fee"`
	GatewayTransferDateToCommerceDate time.Time   `json:"gateway_transfer_date" bson:"gateway_transfer_date_to_commerce_date"`
}

type FlowPaymentStatus struct {
	FlowOrder     uint                       `json:"flowOrder" bson:"flow_order"`
	CommerceOrder string                     `json:"commerceOrder" bson:"commerce_order"`
	RequestDate   string                     `json:"requestDate" bson:"request_date"`
	Status        uint                       `json:"status" bson:"status"`
	Subject       string                     `json:"subject" bson:"subject"`
	Currency      string                     `json:"currency" bson:"currency"`
	Amount        string                     `json:"amount" bson:"amount"`
	Payer         string                     `json:"payer" bson:"payer"`
	Optional      models.GatewayOptionalData `json:"optional" bson:"optional"`
	PendingInfo   FlowPendingInfo            `json:"pendingInfo" bson:"pending_info"`
	PaymentData   FlowPaymentData            `json:"paymentData" bson:"payment_data"`
	MerchantID    string                     `json:"merchantId" bson:"merchant_id"`
}

type FlowPendingInfo struct {
	Media string `json:"media"`
	Date  string `json:"date"`
}

type FlowPaymentData struct {
	Date           string  `json:"date" bson:"date"`
	Media          string  `json:"media" bson:"media"`
	ConversionDate string  `json:"conversionDate" bson:"conversion_date"`
	ConversionRate float64 `json:"conversionRate" bson:"conversion_rate"`
	Amount         string  `json:"amount" bson:"amount"`
	Currency       string  `json:"currency" bson:"currency"`
	Fee            string  `json:"fee" bson:"fee"`
	Balance        float64 `json:"balance" bson:"balance"`
	TransferDate   string  `json:"transferDate" bson:"transfer_date"`
}
