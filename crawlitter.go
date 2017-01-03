package main

const (
	screen_name     = "katyperry" // The most popular Twitter user
	collection_name = "user_graphs"
)

func main() {
	api := CreateTwitterApi()
	user, _ := api.GetUsersShow(screen_name, nil)

	CollectUserGraph(api, user.Id)
}
