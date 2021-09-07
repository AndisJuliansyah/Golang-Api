package config

import "github.com/tkanos/gonfig"

type Configuration struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string
	DB_NAME     string
	DB_PORT     string
}

func Getconfig() Configuration {
	conf := Configuration{}

	gonfig.GetConf("config/config.json", &conf)

	return conf
}
