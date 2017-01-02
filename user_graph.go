package main

import (
	"time"

	"github.com/ChimeraCoder/anaconda"
)

type UserGraph struct {
	User_id   int64 ",omitempty"
	Following []int64
	Followers []int64
	Got_at    time.Time
}

func GetGraphById(api *anaconda.TwitterApi, id int64) (userGraph UserGraph) {
	following, _ := api.GetFriendsUser(id, nil)
	followers, _ := api.GetFollowersUser(id, nil)
	return UserGraph{id, following.Ids, followers.Ids, time.Now()}
}
