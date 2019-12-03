package repository

import (
	"github.com/go-pg/pg/v9"
	"gitlab.com/km/go-kafka-playground/models"
	"gitlab.com/km/go-kafka-playground/service/user"
)

type psqlUserRepository struct {
	db *pg.DB
}

func NewPsqlUserRepository(dbcon *pg.DB) user.PsqlUserRepositoryInf {
	return &psqlUserRepository{
		db: dbcon,
	}
}

func (p *psqlUserRepository) Create(user *models.User) error {
	return p.db.Insert(user)
}
