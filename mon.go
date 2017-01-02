package main

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func GetSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		log.Panic(err)
	}
	return session
}

func ExistsUserID(collection mgo.Collection, user_id int64) bool {
	n, err := collection.Find(bson.M{"user_id": user_id}).Count()
	if err != nil {
		log.Fatal(err)
	}
	return n > 0
}
