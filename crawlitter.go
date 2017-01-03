package main

const screen_name = "katyperry" // The most popular Twitter user

func main() {
	userGraphCollector := NewUserGraphCollector()

	// Get first user_id for graph search
	user, _ := userGraphCollector.Api.GetUsersShow(screen_name, nil)

	userGraphCollector.Collect(user.Id)
}
