package usecase

import (
	"gitlab.com/km/go-kafka-playground/service/kafka"
	"gitlab.com/km/go-kafka-playground/service/user"
)

type kafkaEventUsecase struct {
	userUs user.UserUsecaseInf
}

func NewKafkaEventUsecase(us user.UserUsecaseInf) kafka.KafkaEventUsecase {
	return &kafkaEventUsecase{
		userUs: us,
	}
}
