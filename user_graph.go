package main

import (
	"log"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/ChimeraCoder/anaconda"
)

type UserGraph struct {
	User_id   int64 ",omitempty"
	Following []int64
	Followers []int64
	Got_at    time.Time
}

func getGraphById(api *anaconda.TwitterApi, id int64) (userGraph UserGraph) {
	following, _ := api.GetFriendsUser(id, nil)
	followers, _ := api.GetFollowersUser(id, nil)
	return UserGraph{id, following.Ids, followers.Ids, time.Now()}
}

func existsUserID(collection mgo.Collection, user_id int64) bool {
	n, err := collection.Find(bson.M{"user_id": user_id}).Count()
	if err != nil {
		log.Fatal(err)
	}
	return n > 0
}

func CollectUserGraph(api *anaconda.TwitterApi, user_id int64) {
	mongo := NewMongo()
	session := mongo.GetSession()
	defer session.Close()

	collection := session.DB("crawlitter").C(collection_name)

	if existsUserID(*collection, user_id) {
		log.Print("skip")
	} else {
		userGraph := getGraphById(api, user_id)
		log.Print(userGraph)

		err := collection.Insert(&userGraph)
		if err != nil {
			log.Fatal(err)
		}
	}
}
