package config

import (
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/slovak-egov/einvoice/environment"
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

type Configuration struct {
	Port            int
	RedisUrl        string
	ApiServerUrl    string
	SlovenskoSk     SlovenskoSkConfiguration
	Db              DbConfiguration
	TokenExpiration time.Duration
	LogLevel        log.Level
}

func (c *Configuration) initDb() {
	c.Db = DbConfiguration{
		Host:     environment.Getenv("DB_HOST", c.Db.Host),
		Port:     environment.ParseInt("DB_PORT", c.Db.Port),
		Name:     environment.Getenv("DB_NAME", c.Db.Name),
		User:     environment.Getenv("DB_USER", c.Db.User),
		Password: environment.RequireVar("DB_PASSWORD"),
	}
}

func Init() Configuration{
	authproxyEnv := environment.RequireVar("AUTHPROXY_ENV")
	var config Configuration
	switch authproxyEnv {
	case "prod":
		config = prodConfig
	case "dev":
		config = devConfig
	case "test":
		config = testConfig
	default:
		log.WithField("environment", authproxyEnv).Fatal("config.environment.unknown")
	}

	log.SetFormatter(&log.JSONFormatter{})
	var err error
	logLevel := environment.Getenv("LOG_LEVEL", config.LogLevel.String())
	config.LogLevel, err = log.ParseLevel(logLevel)
	if err != nil {
		log.WithField("log_level", logLevel).Fatal("config.log_level.unknown")
	}

	config.Port = environment.ParseInt("PORT", config.Port)
	config.RedisUrl = environment.Getenv("REDIS_URL", config.RedisUrl)
	config.ApiServerUrl = environment.Getenv("APISERVER_URL", config.ApiServerUrl)

	config.SlovenskoSk = SlovenskoSkConfiguration{
		Url:                environment.Getenv("SLOVENSKO_SK_URL", config.SlovenskoSk.Url),
		ApiTokenPrivateKey: environment.Getenv("API_TOKEN_PRIVATE", config.SlovenskoSk.ApiTokenPrivateKey),
		OboTokenPublicKey:  environment.Getenv("OBO_TOKEN_PUBLIC", config.SlovenskoSk.OboTokenPublicKey),
	}

	config.initDb()

	if tokenExpirationSeconds := environment.ParseInt("TOKEN_EXPIRATION_SECONDS", -1); tokenExpirationSeconds != -1 {
		config.TokenExpiration = time.Duration(tokenExpirationSeconds) * time.Second
	}

	log.Info("config.loaded")

	return config
}
