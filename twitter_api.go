package main

import (
	"github.com/BurntSushi/toml"
	"github.com/ChimeraCoder/anaconda"
)

type twitterConf struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

func CreateTwitterApi() (api *anaconda.TwitterApi) {
	var conf twitterConf
	_, err := toml.DecodeFile("config.tml", &conf)
	if err != nil {
		panic(err)
	}

	anaconda.SetConsumerKey(conf.ConsumerKey)
	anaconda.SetConsumerSecret(conf.ConsumerSecret)
	return anaconda.NewTwitterApi(conf.AccessToken, conf.AccessTokenSecret)
}
