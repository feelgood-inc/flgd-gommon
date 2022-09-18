package kafka

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/avast/retry-go"
	"github.com/feelgood-inc/flgd-gommon/logger"
	"github.com/feelgood-inc/flgd-gommon/models"
	"github.com/segmentio/kafka-go"
)

type WriterConfig struct {
	Logger logger.Logger
}

func NewWriter(writer kafka.Writer, log logger.Logger) *kafka.Writer {
	w := &writer

	if log != nil {
		w.Logger = kafka.LoggerFunc(log.Debugf)
		w.ErrorLogger = kafka.LoggerFunc(log.Errorf)
	}
	return w
}

func WriteToTopicInBackground(
	ctx context.Context,
	writer *kafka.Writer,
	topic string,
	key string,
	value models.KafkaMessage,
	log logger.Logger,
) error {
	if writer == nil {
		log.Errorf("writer is nil")
		return errors.New("writer is nil")
	}

	messageAsBytes, err := json.Marshal(value)
	if err != nil {
		log.Errorf("Error marshalling message: %v", err)
		return err
	}

	go func(cx context.Context, msg []byte, topic string, key string) {
		if err2 := retry.Do(func() error {
			err2 := writer.WriteMessages(ctx, kafka.Message{
				Key:   []byte(key),
				Value: msg,
				Topic: topic,
			})
			if err2 != nil {
				log.Errorf("Error writing message to topic: %v", err2)
				return err2
			}
			return nil
		},
			retry.Attempts(3),
			retry.DelayType(retry.BackOffDelay),
			retry.Context(cx),
		); err2 != nil {
			log.Errorf("Error writing message to topic: %v", err2)
		}
	}(ctx, messageAsBytes, topic, key)

	return nil
}
