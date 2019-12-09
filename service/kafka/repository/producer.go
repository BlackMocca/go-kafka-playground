package repository

import (
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
