package repository

import (
	"time"

	"github.com/Shopify/sarama"
	"gitlab.com/km/go-kafka-playground/config"
	"gitlab.com/km/go-kafka-playground/service/kafka"
)

type kafkaProducerRepository struct {
	customProducer *config.KafkaProducer
}

func NewKafkaProducerRepository(producer *config.KafkaProducer) kafka.KafkaProducerRepository {
	return &kafkaProducerRepository{
		customProducer: producer,
	}
}

func (k *kafkaProducerRepository) PrepareMessage(topic, message string) *sarama.ProducerMessage {
	now := time.Now()
	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: -1,
		Value:     sarama.StringEncoder(message),
		Timestamp: now,
	}

	return msg
}

func (k *kafkaProducerRepository) SendOneMessageWithSync(msg *sarama.ProducerMessage) (int32, int64, error) {
	return k.customProducer.GetSyncProducer().SendMessage(msg)
}
