package models

type Bank struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

type BankAccount struct {
	Bank    Bank   `json:"bank"`
	Number  string `json:"number"`
	Type    string `json:"type"`
	Country string `json:"country"`
	UID     string `json:"uid"`
}
