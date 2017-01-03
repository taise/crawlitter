package main

import (
	"log"

	"github.com/BurntSushi/toml"

	"gopkg.in/mgo.v2"
)

const Database = "crawlitter"

type mongoConf struct {
	Host string
}

type Mongo struct {
	Host     string
	Database string
}

func (self *Mongo) initialize() {
	var conf mongoConf
	_, err := toml.DecodeFile("database.tml", &conf)
	if err != nil {
		log.Panic(err)
	}

	self.Host = conf.Host
	self.Database = Database
}

func NewMongo() *Mongo {
	mongo := &Mongo{}
	mongo.initialize()
	return mongo
}

func (self *Mongo) buildUri() string {
	return "mongodb://" + self.Host
}

func (self *Mongo) GetSession() *mgo.Session {
	session, err := mgo.Dial(self.buildUri())
	if err != nil {
		log.Panic(err)
	}
	return session
}
