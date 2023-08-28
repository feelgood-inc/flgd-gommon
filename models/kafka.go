package models

import "time"

type KafkaMessage struct {
	Metadata *KafkaMetadata `json:"metadata"`
	Data     interface{}    `json:"data"`
	Code     string         `json:"code"`
	Error    *KafkaError    `json:"error"`
}

type KafkaMetadata struct {
	IssuedBy string `json:"issued_by"`
}

type KafkaErrorMessage struct {
	Offset    int64     `json:"offset"`
	Error     string    `json:"error"`
	Time      time.Time `json:"time"`
	Partition int       `json:"partition"`
	Topic     string    `json:"topic"`
}

type KafkaError struct {
	Offset    int64     `json:"offset"`
	Error     string    `json:"error"`
	Time      time.Time `json:"time"`
	Partition int       `json:"partition"`
	Topic     string    `json:"topic"`
}
