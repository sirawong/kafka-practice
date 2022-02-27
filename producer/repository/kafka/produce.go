package kafka

import (
	"encoding/json"
	"events"
	"log"

	"github.com/Shopify/sarama"
)

func (obj kafkaRepo) Produce(event events.Event) error {
	value, err := json.Marshal(event)
	if err != nil {
		return err
	}

	msg := sarama.ProducerMessage{
		Topic: obj.config.MessageTopic,
		Value: sarama.ByteEncoder(value),
	}

	partition, offset, err := obj.producer.SendMessage(&msg)
	if err != nil {
		return err
	}
	log.Println(partition, offset)

	return nil
}
