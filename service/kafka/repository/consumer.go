package repository

import (
	"gitlab.com/km/go-kafka-playground/config"
	"gitlab.com/km/go-kafka-playground/service/kafka"
)

type kafkaConsumerRepository struct {
	customConsumer *config.KafkaConsumer
}

func NewKafkaConsumerRepository(producer *config.KafkaConsumer) kafka.KafkaConsumerRepository {
	return &kafkaConsumerRepository{
		customConsumer: producer,
	}
}
