package models

type KafkaMessage struct {
	Metadata *KafkaMetadata `json:"metadata"`
	Data     interface{}    `json:"data"`
}

type KafkaMetadata struct {
	IssuedBy string `json:"issued_by"`
}
