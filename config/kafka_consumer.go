package config

import (
	"log"

	"github.com/Shopify/sarama"
)

type KafkaConsumer struct {
	consumer sarama.Consumer
}

func NewKafkaConsumer(client sarama.Client) *KafkaConsumer {
	consumerInf, err := sarama.NewConsumerFromClient(client)
	if err != nil {
		log.Fatal(err)
	}

	return &KafkaConsumer{
		consumer: consumerInf,
	}
}

func (k KafkaConsumer) GetConsumer() sarama.Consumer {
	return k.consumer
}
