package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/ChimeraCoder/anaconda"
)

type Config struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

func main() {
	var conf Config
	if _, err := toml.DecodeFile("config.tml", &conf); err != nil {
		panic(err)
	}

	anaconda.SetConsumerKey(conf.ConsumerKey)
	anaconda.SetConsumerSecret(conf.ConsumerSecret)
	api := anaconda.NewTwitterApi(conf.AccessToken, conf.AccessTokenSecret)

	user, _ := api.GetUsersShow("_eurk", nil)
	fmt.Println(user)
}
