package config

import (
	log "github.com/sirupsen/logrus"

	"github.com/slovak-egov/einvoice/environment"
	"time"
)

type DbConfiguration struct {
	Host     string
	Port     int
	Name     string
	User     string
	Password string
}

type SlovenskoSkConfiguration struct {
	Url                string
	ApiTokenPrivateKey string
	OboTokenPublicKey  string
}

type configuration struct {
	Port            int
	RedisUrl        string
	ApiServerUrl    string
	SlovenskoSk     SlovenskoSkConfiguration
	Db              DbConfiguration
	TokenExpiration time.Duration
	LogLevel        log.Level
}

var Config = configuration{}

func (c *configuration) initDb() {
	c.Db = DbConfiguration{
		Host:     environment.Getenv("DB_HOST", c.Db.Host),
		Port:     environment.ParseInt("DB_PORT", c.Db.Port),
		Name:     environment.Getenv("DB_NAME", c.Db.Name),
		User:     environment.Getenv("DB_USER", c.Db.User),
		Password: environment.RequireVar("DB_PASSWORD"),
	}
}

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
		Url:                environment.Getenv("SLOVENSKO_SK_URL", Config.SlovenskoSk.Url),
		ApiTokenPrivateKey: environment.RequireVar("API_TOKEN_PRIVATE"),
		OboTokenPublicKey:  environment.RequireVar("OBO_TOKEN_PUBLIC"),
	}

	Config.initDb()

	if tokenExpirationSeconds := environment.ParseInt("TOKEN_EXPIRATION_SECONDS", -1); tokenExpirationSeconds != -1 {
		Config.TokenExpiration = time.Duration(tokenExpirationSeconds) * time.Second
	}

	log.Info("config.loaded")
}
