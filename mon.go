package main

import (
	"log"

	"gopkg.in/mgo.v2"
)

func GetSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		log.Panic(err)
	}
	return session
}
