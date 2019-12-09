package usecase

import (
	"gitlab.com/km/go-kafka-playground/service/kafka"
)

type kafkaUsecase struct {
	producerRepo kafka.KafkaProducerRepository
	consumerRepo kafka.KafkaConsumerRepository
}

func NewKafkaUsecase(producer kafka.KafkaProducerRepository, consumer kafka.KafkaConsumerRepository) kafka.KafkaUsecase {
	return &kafkaUsecase{
		producerRepo: producer,
		consumerRepo: consumer,
	}
}
