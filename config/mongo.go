package config

import (
	"log"
	"time"

	mgo "github.com/globalsign/mgo"
)

/*
	define db
*/
var (
	DB_APPEXAMPLE = "app_example"
)

/*
	define collection
*/
var (
	APP_EXAMPLE_COLLECT = "users"
)

func initCollection(session *mgo.Session) {
	col := session.DB(DB_APPEXAMPLE).C(APP_EXAMPLE_COLLECT)
	colInfo := mgo.CollectionInfo{
		DisableIdIndex: false,
		ForceIdIndex:   true,
	}

	if err := col.Create(&colInfo); err != nil {
		log.Println(err)
	}
}

func NewMongoSession() *mgo.Session {
	mongoURL := GetEnv("MONGO_DATABASE_URL", "mongodb:://mongoadmin:mongoadmin@mongo_db:27017/"+DB_APPEXAMPLE)
	timeout := 60 * time.Second

	session, err := mgo.DialWithTimeout(mongoURL, timeout)
	if err != nil {
		log.Fatal(err)
	}

	session.SetBatch(10000)
	initCollection(session)
	return session

}
