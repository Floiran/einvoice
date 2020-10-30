package config

import (
	log "github.com/sirupsen/logrus"

	"github.com/slovak-egov/einvoice/environment"
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
	LogLevel     log.Level
}

var Config = configuration{}

func InitConfig() {
	authproxyEnv := environment.RequireVar("AUTHPROXY_ENV")

	switch authproxyEnv {
	case "prod":
		Config = prodConfig
	case "dev":
		Config = devConfig
	default:
		log.WithField("environment", authproxyEnv).Fatal("config.environment.unknown")
	}

	log.SetFormatter(&log.JSONFormatter{})
	var err error
	logLevel := environment.Getenv("LOG_LEVEL", Config.LogLevel.String())
	Config.LogLevel, err = log.ParseLevel(logLevel)
	if err != nil {
		log.WithField("log_level", logLevel).Fatal("config.log_level.unknown")
	}

	Config.Port = environment.ParseInt("PORT", Config.Port)
	Config.RedisUrl = environment.Getenv("REDIS_URL", Config.RedisUrl)
	Config.ApiServerUrl = environment.Getenv("APISERVER_URL", Config.ApiServerUrl)

	Config.SlovenskoSk = SlovenskoSkConfiguration{
		Url: environment.Getenv("SLOVENSKO_SK_URL", Config.SlovenskoSk.Url),
		ApiTokenPrivateKey: environment.RequireVar("API_TOKEN_PRIVATE"),
		OboTokenPublicKey: environment.RequireVar("OBO_TOKEN_PUBLIC"),
	}

	log.Info("Config loaded")
}
