package config

import (
	"log"

	"github.com/slovak-egov/einvoice/environment"
)

type configuration struct {
	DbHost     string
	DbPort     int
	DbName     string
	DbUser     string
	DbPassword string
}

var Config = configuration{}

func InitConfig() {
	Config.DbHost = environment.RequireVar("DB_HOST")
	Config.DbPort = environment.ParseInt("DB_PORT")
	Config.DbName = environment.RequireVar("DB_NAME")
	Config.DbUser = environment.RequireVar("DB_USER")
	Config.DbPassword = environment.RequireVar("DB_PASSWORD")

	log.Println("Config loaded")
}
