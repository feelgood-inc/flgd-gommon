package kafka

import (
	"github.com/feelgood-inc/flgd-gommon/logger"
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
