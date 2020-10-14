package config

import (
	"github.com/slovak-egov/einvoice/environment"
	"log"
)

type SlovenskoSkConfiguration struct {
	Url                string
	ApiTokenPrivateKey string
	OboTokenPublicKey  string
}

type configuration struct {
	Port         int
	RedisUrl     string
	ApiServerUrl string
	SlovenskoSk  SlovenskoSkConfiguration
}

var Config = configuration{}

func InitConfig() {
	Config.Port = environment.ParseInt("PORT")

	Config.RedisUrl = environment.RequireVar("REDIS_URL")
	Config.ApiServerUrl = environment.RequireVar("APISERVER_URL")

	Config.SlovenskoSk.Url = environment.RequireVar("SLOVENSKO_SK_URL")
	Config.SlovenskoSk.ApiTokenPrivateKey = environment.RequireVar("API_TOKEN_PRIVATE")
	Config.SlovenskoSk.OboTokenPublicKey = environment.RequireVar("OBO_TOKEN_PUBLIC")

	log.Println("Config loaded")
}
