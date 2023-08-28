package kafka

import (
	"context"
	"fmt"
	"github.com/feelgood-inc/flgd-gommon/models"
	"github.com/goccy/go-json"
	"github.com/segmentio/kafka-go"
)

func UnmarshalToKafkaMessage(data []byte) (models.KafkaMessage, error) {
	var kafkaMessage models.KafkaMessage
	if err := json.Unmarshal(data, &kafkaMessage); err != nil {
		return kafkaMessage, err
	}

	if kafkaMessage == (models.KafkaMessage{}) || kafkaMessage.Data == nil {
		return kafkaMessage, fmt.Errorf("error unmarshaling to KafkaMessage")
	}

	return kafkaMessage, nil
}

func MarshalKafkaDataToStruct(data []byte, structToMarshal interface{}) error {
	kafkaMessage, err := UnmarshalToKafkaMessage(data)
	if err != nil {
		return err
	}
	if kafkaMessage.Data == nil {
		return fmt.Errorf("error unmarshaling to KafkaMessage")
	}
	if kafkaMessage == (models.KafkaMessage{}) {
		return fmt.Errorf("error unmarshaling to KafkaMessage")
	}

	asBytes, err := json.Marshal(kafkaMessage.Data)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(asBytes, &structToMarshal); err != nil {
		return err
	}

	return nil
}

func PublishToDLQ(ctx context.Context, writer *kafka.Writer, dlqTopic string, message models.KafkaMessage, error string) error {
	message.Error.Error = error

	messageAsBytes, err := json.Marshal(message)
	if err != nil {
		return err
	}

	if err := writer.WriteMessages(ctx, kafka.Message{
		Value: messageAsBytes,
		Topic: dlqTopic,
	}); err != nil {
		return err
	}

	return nil
}
