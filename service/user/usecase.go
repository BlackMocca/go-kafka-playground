package user

import "gitlab.com/km/go-kafka-playground/models"

type UserUsecaseInf interface {
	Create(user *models.User) error
	CreateIntoMongoDB(user *models.User) error
	FetchOne(id int) (*models.User, error)
	InvokeCreateEvent(user *models.User) (int32, int64, error)
}
