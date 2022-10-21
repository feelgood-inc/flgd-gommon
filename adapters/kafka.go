package adapters

import (
	"context"
	"crypto/tls"
	"github.com/feelgood-inc/flgd-gommon/config"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
)

type KafkaConnectionConfig struct {
	Brokers []string
}

type KafkaMechanismConfig struct {
	Username string
	Password string
}

type KafkaWriterWithMechanismConfig struct {
	MechanismConfig KafkaMechanismConfig
	Addr            []string
	Topic           string
	Balancer        kafka.Balancer
	Async           bool
	MaxAttempts     int
}

func NewKafkaConn(cfg *config.Config) (*kafka.Conn, error) {
	return kafka.DialContext(context.Background(), "tcp", cfg.Kafka.Brokers[0])
}

func NewWriterWithMechanism(cfg *KafkaWriterWithMechanismConfig) *kafka.Writer {
	mechanism, err := scram.Mechanism(scram.SHA512, cfg.MechanismConfig.Username, cfg.MechanismConfig.Password)
	if err != nil {
		panic(err)
	}

	transport := &kafka.Transport{
		SASL: mechanism,
		TLS:  &tls.Config{},
	}

	return &kafka.Writer{
		Addr:        kafka.TCP(cfg.Addr...),
		Topic:       cfg.Topic,
		Balancer:    cfg.Balancer,
		Async:       cfg.Async,
		MaxAttempts: cfg.MaxAttempts,
		Transport:   transport,
	}
}

func NewReaderWithTLS(cfg *KafkaConnectionConfig) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:   cfg.Brokers,
		Topic:     "test",
		Partition: 0,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})
}
