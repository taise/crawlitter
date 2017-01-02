package main

import "log"

func main() {
	api := CreateTwitterApi()
	user, _ := api.GetUsersShow("_eurk", nil)

	session := GetSession()
	defer session.Close()

	var db string = "crawlitter"
	var collection_name string = "user_graphs"

	collection := session.DB(db).C(collection_name)

	if ExistsUserID(*collection, user.Id) {
		log.Print("skip")
	} else {
		userGraph := GetGraphById(api, user.Id)
		log.Print(userGraph)

		err := collection.Insert(&userGraph)
		if err != nil {
			log.Fatal(err)
		}
	}
}
