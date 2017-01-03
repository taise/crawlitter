package main

import "log"

const (
	screen_name     = "katyperry" // The most popular Twitter user
	collection_name = "user_graphs"
)

func main() {
	api := CreateTwitterApi()
	user, _ := api.GetUsersShow(screen_name, nil)

	collection := GetCollection(collection_name)

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
