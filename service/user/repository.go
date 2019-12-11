package user

import "gitlab.com/km/go-kafka-playground/models"

type PsqlUserRepositoryInf interface {
	FetchOne(id int) (*models.User, error)
	Create(user *models.User) error
}

type MongoUserRepositoryInf interface {
	Create(user *models.User) error
}
