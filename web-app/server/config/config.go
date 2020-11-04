package config

import (
	log "github.com/sirupsen/logrus"

	"github.com/slovak-egov/einvoice/environment"
)

type Configuration struct {
	Port                int
	AuthServerUrl       string
	ClientBuildDir      string
	SlovenskoSkLoginUrl string
	LogLevel            log.Level
}

func Init() Configuration {
	webserverEnv := environment.RequireVar("WEBSERVER_ENV")
	var config Configuration
	switch webserverEnv {
	case "prod":
		config = prodConfig
	case "dev":
		config = devConfig
	default:
		log.WithField("environment", webserverEnv).Fatal("config.environment.unknown")
	}

	log.SetFormatter(&log.JSONFormatter{})
	var err error
	logLevel := environment.Getenv("LOG_LEVEL", config.LogLevel.String())
	config.LogLevel, err = log.ParseLevel(logLevel)
	if err != nil {
		log.WithField("log_level", logLevel).Fatal("config.log_level.unknown")
	}

	config.Port = environment.ParseInt("PORT", config.Port)
	config.AuthServerUrl = environment.Getenv("AUTH_SERVER_URL", config.AuthServerUrl)
	config.ClientBuildDir = environment.Getenv("CLIENT_BUILD_DIR", config.ClientBuildDir)
	config.SlovenskoSkLoginUrl = environment.Getenv("SLOVENSKO_SK_LOGIN_URL", config.SlovenskoSkLoginUrl)

	log.Info("config.loaded")
	return config
}
