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

func (k kafkaEventUsecase) CreateUserIntoMongo(id int) error {
	user, err := k.userUs.FetchOne(id)
	if err != nil {
		return err
	}

	if err = k.userUs.CreateIntoMongoDB(user); err != nil {
		return err
	}
	return nil
}
