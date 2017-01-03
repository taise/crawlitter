package main

import (
	"log"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/ChimeraCoder/anaconda"
)

const collection_name = "user_graphs"

type UserGraph struct {
	User_id   int64 ",omitempty"
	Following []int64
	Followers []int64
	Got_at    time.Time
}

type UserGraphCollector struct {
	Api        *anaconda.TwitterApi
	Collection *mgo.Collection
}

func (self *UserGraphCollector) initialize() {
	self.Api = CreateTwitterApi()
	mongo := NewMongo()
	self.Collection = mongo.GetSession().DB(mongo.DB).C(collection_name)
}

func NewUserGraphCollector() *UserGraphCollector {
	userGraphCollector := &UserGraphCollector{}
	userGraphCollector.initialize()
	return userGraphCollector
}

func (self *UserGraphCollector) existsUserID(user_id int64) bool {
	n, err := self.Collection.Find(bson.M{"user_id": user_id}).Count()
	if err != nil {
		log.Fatal(err)
	}
	return n > 0
}

func (self *UserGraphCollector) breadthFirstSearch(ids []int64) int {
	return 1
}

func (self *UserGraphCollector) getGraphById(id int64) (userGraph UserGraph) {
	following, _ := self.Api.GetFriendsUser(id, nil)
	followers, _ := self.Api.GetFollowersUser(id, nil)
	return UserGraph{id, following.Ids, followers.Ids, time.Now()}
}

func (self *UserGraphCollector) Collect(user_id int64) {

	if self.existsUserID(user_id) {
		log.Print("skip")
	} else {
		userGraph := self.getGraphById(user_id)
		log.Print(userGraph)

		err := self.Collection.Insert(&userGraph)
		if err != nil {
			log.Fatal(err)
		}
	}
}