package repository

import (
	"gitlab.com/km/go-kafka-playground/config"
	"gitlab.com/km/go-kafka-playground/service/kafka"
)

type kafkaConsumerRepository struct {
	customConsumer *config.KafkaConsumerGroup
}

func NewKafkaConsumerRepository(producer *config.KafkaConsumerGroup) kafka.KafkaConsumerRepository {
	return &kafkaConsumerRepository{
		customConsumer: producer,
	}
}
