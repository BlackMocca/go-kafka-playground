package models

import (
	"reflect"
	"time"
)

var (
	TopicUser = "users"
)

type User struct {
	TableName struct{} `json:"-" pg:"users" bson:"-"`

	ID        int        `json:"id" pg:"id" bson:"id"`
	Email     string     `json:"email" pg:"email" bson:"email"`
	FirstName string     `json:"firstname" pg:"firstname" bson:"firstname"`
	LastName  string     `json:"lastname" pg:"lastname" bson:"lastname"`
	Age       int        `json:"age" pg:"age" bson:"age"`
	CreatedAt *time.Time `json:"created_at" pg:"created_at" bson:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" pg:"updated_at" bson:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" pg:"deleted_at" bson:"deleted_at"`
}

func NewUserWithParams(params map[string]interface{}, user *User) *User {
	if user != nil {
		user = new(User)
	}

	if v, ok := params["id"]; ok {
		user.ID = v.(int)
	}

	if v, ok := params["email"]; ok {
		user.Email = v.(string)
	}

	if v, ok := params["firstname"]; ok {
		user.FirstName = v.(string)
	}
	if v, ok := params["lastname"]; ok {
		user.LastName = v.(string)
	}
	if v, ok := params["age"]; ok {
		user.Age = v.(int)
	}

	if v, ok := params["created_at"]; ok {
		if reflect.ValueOf(v).Kind() == reflect.String {
			t, _ := time.Parse(time.RFC3339, v.(string))
			user.CreatedAt = &t
		} else {
			t := v.(time.Time)
			user.CreatedAt = &t
		}
	}

	if v, ok := params["updated_at"]; ok {
		if reflect.ValueOf(v).Kind() == reflect.String {
			t, _ := time.Parse(time.RFC3339, v.(string))
			user.UpdatedAt = &t
		} else {
			t := v.(time.Time)
			user.UpdatedAt = &t
		}
	}

	if v, ok := params["deleted_at"]; ok {
		if reflect.ValueOf(v).Kind() == reflect.String {
			t, _ := time.Parse(time.RFC3339, v.(string))
			user.DeletedAt = &t
		} else {
			t := v.(time.Time)
			user.DeletedAt = &t
		}
	}

	return user
}
