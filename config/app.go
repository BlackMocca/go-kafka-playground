package config

import (
	"os"
	"reflect"

	"github.com/Shopify/sarama"
	mgo "github.com/globalsign/mgo"
	"github.com/go-pg/pg/v9"
	"github.com/joho/godotenv"
)

type Config struct {
	PGORM *pg.DB
	MONGO *mgo.Session
	Kafka sarama.Client
}

func init() {
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}
}

func NewConfigWithService(args ...interface{}) *Config {
	c := Config{}
	for i, _ := range args {
		service := reflect.TypeOf(args[i]).String()

		switch service {
		case "*mgo.Session":
			c.MONGO = args[i].(*mgo.Session)
		case "*pg.DB":
			c.PGORM = args[i].(*pg.DB)
		}
	}
	return &c
}

func (c *Config) SetService(arg interface{}) {
	service := reflect.TypeOf(arg).String()
	switch service {
	case "*mgo.Session":
		c.MONGO = arg.(*mgo.Session)
	case "*pg.DB":
		c.PGORM = arg.(*pg.DB)
	}
}

func GetEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
