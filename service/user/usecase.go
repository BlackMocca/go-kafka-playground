package user

import "gitlab.com/km/go-kafka-playground/models"

type UserUsecaseInf interface {
	Create(user *models.User) error
}
