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

func (k *kafkaUsecase) SendMessage(topic, message string) (int32, int64, error) {
	msg := k.producerRepo.PrepareMessage(topic, message)
	partition, offset, err := k.producerRepo.SendOneMessageWithSync(msg)
	if err != nil {
		return partition, offset, err
	}
	return partition, offset, nil
}
