package config

import (
	log "github.com/sirupsen/logrus"

	"github.com/slovak-egov/einvoice/environment"
)

type Configuration struct {
	DbHost     string
	DbPort     int
	DbName     string
	DbUser     string
	DbPassword string
}

func Init() Configuration {
	config := Configuration{
		DbHost:     environment.Getenv("DB_HOST", "localhost"),
		DbPort:     environment.ParseInt("DB_PORT", 5432),
		DbName:     environment.Getenv("DB_NAME", "einvoice"),
		DbUser:     environment.Getenv("DB_USER", "postgres"),
		DbPassword: environment.RequireVar("DB_PASSWORD"),
	}

	log.Info("config.loaded")

	return config
}
