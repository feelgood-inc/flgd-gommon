package kafka

import (
	"github.com/feelgood-inc/flgd-gommon/logger"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/compress"
)

type WriterConfig struct {
	Brokers     []string
	Compression compress.Compression
	Logger      logger.Logger
	Balancer    kafka.Balancer
}

func NewWriter(writer kafka.Writer, config WriterConfig) *kafka.Writer {
	w := &kafka.Writer{
		Addr:         kafka.TCP(config.Brokers...),
		Topic:        writer.Topic,
		Balancer:     config.Balancer,
		RequiredAcks: writer.RequiredAcks,
		MaxAttempts:  writer.MaxAttempts,
		Logger:       kafka.LoggerFunc(config.Logger.Debugf),
		ErrorLogger:  kafka.LoggerFunc(config.Logger.Errorf),
		Compression:  config.Compression,
		ReadTimeout:  writer.ReadTimeout,
		WriteTimeout: writer.WriteTimeout,
	}
	return w
}
