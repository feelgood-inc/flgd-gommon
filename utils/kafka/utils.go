package kafka

import (
	"encoding/json"
	"fmt"
	"github.com/feelgood-inc/flgd-gommon/models"
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
