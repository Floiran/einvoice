package config

import (
	"github.com/slovak-egov/einvoice/environment"
	"log"
)

type configuration struct {
	Port                int
	AuthServerUrl       string
	ClientBuildDir      string
	SlovenskoSkLoginUrl string
}

var Config = configuration{}

func InitConfig() {
	Config.Port = environment.ParseInt("PORT")
	Config.AuthServerUrl = environment.RequireVar("AUTH_SERVER_URL")
	Config.ClientBuildDir = environment.RequireVar("CLIENT_BUILD_DIR")
	Config.SlovenskoSkLoginUrl = environment.RequireVar("SLOVENSKO_SK_LOGIN_URL")

	log.Println("Config loaded")
}
