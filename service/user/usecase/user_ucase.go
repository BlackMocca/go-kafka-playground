package usecase

import (
	"strconv"

	"gitlab.com/km/go-kafka-playground/models"
	"gitlab.com/km/go-kafka-playground/service/kafka"
	"gitlab.com/km/go-kafka-playground/service/user"
)

type userUsecase struct {
	psqlUserRepo  user.PsqlUserRepositoryInf
	mongoUserRepo user.MongoUserRepositoryInf
	kafkaUs       kafka.KafkaUsecase
}

func NewUserUsecase(uRepo user.PsqlUserRepositoryInf, mRepo user.MongoUserRepositoryInf, kUs kafka.KafkaUsecase) user.UserUsecaseInf {
	return &userUsecase{
		psqlUserRepo:  uRepo,
		mongoUserRepo: mRepo,
		kafkaUs:       kUs,
	}
}

func (u *userUsecase) InvokeCreateEvent(user *models.User) (int32, int64, error) {
	topic := models.TopicUser
	message := strconv.Itoa(user.ID)
	partition, offset, err := u.kafkaUs.SendMessage(topic, message)
	if err != nil {
		return 0, 0, err
	}

	return partition, offset, nil
}

func (u *userUsecase) Create(user *models.User) error {
	return u.psqlUserRepo.Create(user)
}

func (u *userUsecase) CreateIntoMongoDB(user *models.User) error {
	return u.mongoUserRepo.Create(user)
}

func (u *userUsecase) FetchOne(id int) (*models.User, error) {
	return u.psqlUserRepo.FetchOne(id)
}
