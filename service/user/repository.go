package user

import "gitlab.com/km/go-kafka-playground/models"

type PsqlUserRepositoryInf interface {
	Create(user *models.User) error
}
