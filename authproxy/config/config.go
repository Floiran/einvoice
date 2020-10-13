package config

import (
	"log"
	"github.com/slovak-egov/einvoice/environment"
)

type configuration struct {
	Port int
	RedisUrl string
	ApiServerUrl string
}

var Config = configuration{}

func InitConfig() {
	Config.Port = environment.ParseInt("PORT")

	Config.RedisUrl = environment.RequireVar("REDIS_URL")
	Config.ApiServerUrl = environment.RequireVar("APISERVER_URL")

	log.Println("Config loaded")
}
