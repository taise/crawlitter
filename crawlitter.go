package main

import "fmt"

func main() {

	api := CreateTwitterApi()
	user, _ := api.GetUsersShow("_eurk", nil)

	userGraph := GetGraphById(api, user.Id)
	fmt.Println(userGraph)
}
