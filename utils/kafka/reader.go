package kafka

import (
	"github.com/feelgood-inc/flgd-gommon/logger"
	"github.com/segmentio/kafka-go"
)

type Reader struct {
	log logger.Logger
}

func NewReader(cfg kafka.ReaderConfig, log logger.Logger) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:                cfg.Brokers,
		GroupID:                cfg.GroupID,
		Topic:                  cfg.Topic,
		MinBytes:               cfg.MinBytes,
		MaxBytes:               cfg.MaxBytes,
		QueueCapacity:          cfg.QueueCapacity,
		HeartbeatInterval:      cfg.HeartbeatInterval,
		CommitInterval:         cfg.CommitInterval,
		PartitionWatchInterval: cfg.PartitionWatchInterval,
		Logger:                 kafka.LoggerFunc(log.Debugf),
		ErrorLogger:            kafka.LoggerFunc(log.Errorf),
		MaxAttempts:            cfg.MaxAttempts,
		Dialer: &kafka.Dialer{
			Timeout: cfg.Dialer.Timeout,
		},
	})
}
