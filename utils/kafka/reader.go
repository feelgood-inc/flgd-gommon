package kafka

import (
	"context"
	"crypto/tls"
	"github.com/feelgood-inc/flgd-gommon/logger"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/compress"
	"github.com/segmentio/kafka-go/sasl/scram"
	"github.com/spf13/viper"
	"sync"
	"time"
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

type ConsumerReaderConfig struct {
	Brokers                []string
	MinBytes               int
	MaxBytes               int
	QueueCapacity          int
	HeartbeatInterval      time.Duration
	CommitInterval         time.Duration
	MaxAttempts            int
	DialTimeout            time.Duration
	PartitionWatchInterval time.Duration
}

type ConsumerWriterConfig struct {
	Brokers      []string
	RequiredAcks int
	MaxAttempts  int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func ConsumeFromTopic(
	ctx context.Context,
	cancel context.CancelFunc,
	groupID string,
	topic string,
	workersNum int,
	workerFunc func(ctx context.Context, cancel context.CancelFunc, r *kafka.Reader, w *kafka.Writer, wg *sync.WaitGroup, workerID int),
	readerConfig ConsumerReaderConfig,
	writerConfig ConsumerWriterConfig,
	logger logger.Logger,
) {
	mechanism, err := scram.Mechanism(scram.SHA256, viper.GetString("KAFKA_USERNAME"), viper.GetString("KAFKA_PASSWORD"))
	if err != nil {
		logger.Error("failed to create mechanism", err)
		cancel()
		return
	}
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:                readerConfig.Brokers,
		GroupID:                groupID,
		Topic:                  topic,
		MinBytes:               readerConfig.MinBytes,
		MaxBytes:               readerConfig.MaxBytes,
		QueueCapacity:          readerConfig.QueueCapacity,
		HeartbeatInterval:      readerConfig.HeartbeatInterval,
		CommitInterval:         readerConfig.CommitInterval,
		PartitionWatchInterval: readerConfig.PartitionWatchInterval,
		Logger:                 kafka.LoggerFunc(logger.Debugf),
		ErrorLogger:            kafka.LoggerFunc(logger.Errorf),
		MaxAttempts:            readerConfig.MaxAttempts,
		Dialer: &kafka.Dialer{
			Timeout:       readerConfig.DialTimeout,
			SASLMechanism: mechanism,
			TLS:           &tls.Config{},
		},
	})
	defer cancel()
	defer func() {
		if err := r.Close(); err != nil {
			logger.Errorf("r.Close", err)
			cancel()
		}
	}()

	w := NewWriter(kafka.Writer{
		Addr:         kafka.TCP(writerConfig.Brokers...),
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: kafka.RequiredAcks(writerConfig.RequiredAcks),
		MaxAttempts:  writerConfig.MaxAttempts,
		Logger:       kafka.LoggerFunc(logger.Debugf),
		ErrorLogger:  kafka.LoggerFunc(logger.Errorf),
		Compression:  compress.Snappy,
		ReadTimeout:  writerConfig.ReadTimeout,
		WriteTimeout: writerConfig.WriteTimeout,
	}, logger)
	defer func() {
		if err := w.Close(); err != nil {
			logger.Errorf("w.Close", err)
			cancel()
		}
	}()

	logger.Infof("Starting consumer group: %v", r.Config().GroupID)

	wg := &sync.WaitGroup{}
	for i := 0; i <= workersNum; i++ {
		wg.Add(1)
		go workerFunc(ctx, cancel, r, w, wg, i)
	}
	wg.Wait()
}
