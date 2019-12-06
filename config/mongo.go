package config

import (
	"log"
	"time"

	mgo "github.com/globalsign/mgo"
)

var (
	defaultDBName = "app_example"
)

func NewMongoSession() *mgo.Session {
	mongoURL := GetEnv("MONGO_DATABASE_URL", "mongodb:://mongo:mongo@mongo_db:27017/"+defaultDBName)
	timeout := 60 * time.Second

	session, err := mgo.DialWithTimeout(mongoURL, timeout)
	if err != nil {
		log.Fatal(err)
	}

	session.SetBatch(10000)
	return session

}
