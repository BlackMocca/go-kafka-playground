package models

import (
	"github.com/go-pg/pg"
)

type User struct {
	TableName struct{} `json:"-" pg:"users"`

	ID        int         `json:"id" pg:"id"`
	Email     string      `json:"email" pg:"email"`
	FirstName string      `json:"firstname" pg:"firstname"`
	LastName  string      `json:"lastname" pg:"lastname"`
	Age       int         `json:"age" pg:"age"`
	CreatedAt pg.NullTime `json:"created_at" pg:"created_at"`
	UpdatedAt pg.NullTime `json:"updated_at" pg:"updated_at"`
	DeletedAt pg.NullTime `json:"deleted_at" pg:"deleted_at"`
}
