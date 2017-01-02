package main

import (
	"github.com/BurntSushi/toml"
	"github.com/ChimeraCoder/anaconda"
)

type Config struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

func CreateTwitterApi() (api *anaconda.TwitterApi) {
	var conf Config
	if _, err := toml.DecodeFile("config.tml", &conf); err != nil {
		panic(err)
	}

	anaconda.SetConsumerKey(conf.ConsumerKey)
	anaconda.SetConsumerSecret(conf.ConsumerSecret)
	return anaconda.NewTwitterApi(conf.AccessToken, conf.AccessTokenSecret)
}
