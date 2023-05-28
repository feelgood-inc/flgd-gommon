package models

import "time"

type Invoice struct {
	ID                  string    `json:"id"`
	ExternalID          string    `json:"external_id"`
	Amount              int       `json:"amount"`
	IssuedByTributaryID string    `json:"issued_by_tributary_id"`
	IssuedToTributaryID string    `json:"issued_to_tributary_id"`
	IssuedByUserUID     string    `json:"issued_by_user_uid"`
	IssuedToUserUID     string    `json:"issued_to_user_uid"`
	IssueDate           string    `json:"issue_date"`
	DocumentURL         string    `json:"document_url"`
	ForResourceType     string    `json:"for_resource_type"`
	ForResourceID       string    `json:"for_resource_id"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	CancelledAt         time.Time `json:"cancelled_at"`
	DeletedAt           time.Time `json:"deleted_at"`
}
