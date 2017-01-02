package main

import (
	"log"

	"gopkg.in/mgo.v2"
)

func main() {
	api := CreateTwitterApi()
	user, _ := api.GetUsersShow("_eurk", nil)

	userGraph := GetGraphById(api, user.Id)
	log.Print(userGraph)

	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		log.Panic(err)
	}
	defer session.Close()

	collection := session.DB("crawlitter").C("user_graphs")
	err = collection.Insert(&userGraph)
	if err != nil {
		log.Fatal(err)
	}
}
