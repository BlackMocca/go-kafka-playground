package config

import (
	"log"

	"github.com/Shopify/sarama"
)

type KafkaProducer struct {
	producer sarama.AsyncProducer
}

func NewKafkaProducer(client sarama.Client) *KafkaProducer {
	producerInf, err := sarama.NewAsyncProducerFromClient(client)
	if err != nil {
		log.Fatal(err)
	}

	return &KafkaProducer{
		producer: producerInf,
	}
}

func (k KafkaProducer) GetProducer() sarama.AsyncProducer {
	return k.producer
}
