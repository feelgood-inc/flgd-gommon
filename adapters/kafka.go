package adapters

import (
	"context"
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
