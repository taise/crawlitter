package main

import "log"

func main() {
	api := CreateTwitterApi()
	user, _ := api.GetUsersShow("_eurk", nil)

	userGraph := GetGraphById(api, user.Id)
	log.Print(userGraph)

	session := GetSession()
	defer session.Close()

	var db string = "crawlitter"
	var collection_name string = "user_graphs"

	collection := session.DB(db).C(collection_name)
	err := collection.Insert(&userGraph)
	if err != nil {
		log.Fatal(err)
	}
}
