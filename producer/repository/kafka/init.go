package kafka

import (
	"events"
	"producer/config"

	"github.com/Shopify/sarama"
)

type kafkaRepo struct {
	producer sarama.SyncProducer
	config   *config.Config
}

type Kafka interface {
	Produce(event events.Event) error
}

func New(producer sarama.SyncProducer, config *config.Config) Kafka {
	return kafkaRepo{producer, config}
}
