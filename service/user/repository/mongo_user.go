package repository

import (
	"github.com/globalsign/mgo"
	"gitlab.com/km/go-kafka-playground/models"
	"gitlab.com/km/go-kafka-playground/service/user"
)

type mongoUserRepository struct {
	Conn *mgo.Session
	DB   *mgo.Database
}

func NewMongoUserRepository(Session *mgo.Session, dbName string) user.MongoUserRepositoryInf {
	return &mongoUserRepository{Session, Session.DB(dbName)}
}

func (m *mongoUserRepository) Create(user *models.User) error {
	c := m.DB.C(APP_EXAMPLE_COLLECT)

	return c.Insert(user)
}
