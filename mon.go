package main

import (
	"log"

	"github.com/BurntSushi/toml"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const Database = "crawlitter"

type mongoConf struct {
	Host string
}

func buildMongoUri() string {
	var conf mongoConf
	_, err := toml.DecodeFile("database.tml", &conf)
	if err != nil {
		log.Panic(err)
	}
	return "mongodb://" + conf.Host
}

func getSession() *mgo.Session {
	session, err := mgo.Dial(buildMongoUri())
	if err != nil {
		log.Panic(err)
	}
	return session
}

func GetCollection(name string) *mgo.Collection {
	session := getSession()
	defer session.Close()

	return session.DB(Database).C(name)
}

func ExistsUserID(collection mgo.Collection, user_id int64) bool {
	n, err := collection.Find(bson.M{"user_id": user_id}).Count()
	if err != nil {
		log.Fatal(err)
	}
	return n > 0
}
