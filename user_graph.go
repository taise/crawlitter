package main

import "github.com/ChimeraCoder/anaconda"

type UserGraph struct {
	Id        int64
	following []int64
	followers []int64
}

func GetGraphById(api *anaconda.TwitterApi, id int64) (userGraph UserGraph) {
	following, _ := api.GetFriendsUser(id, nil)
	followers, _ := api.GetFollowersUser(id, nil)
	return UserGraph{id, following.Ids, followers.Ids}
}
