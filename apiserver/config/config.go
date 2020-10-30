package config

import (
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

type Configuration struct {
	Db                   DbConfiguration
	Port                 int
	D16bXsdPath          string
	Ubl21XsdPath         string
	SlowStorageType      string
	LocalStorageBasePath string
	GcsBucket            string
	LogLevel             log.Level
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

func Init() Configuration {
	apiserverEnv := environment.RequireVar("APISERVER_ENV")
	var config Configuration
	switch apiserverEnv {
	case "prod":
		config = prodConfig
	case "dev":
		config = devConfig
	case "test":
		config = testConfig
	default:
		log.WithField("environment", apiserverEnv).Fatal("config.environment.unknown")
	}

	log.SetFormatter(&log.JSONFormatter{})
	var err error
	logLevel := environment.Getenv("LOG_LEVEL", config.LogLevel.String())
	config.LogLevel, err = log.ParseLevel(logLevel)
	if err != nil {
		log.WithField("log_level", logLevel).Fatal("config.log_level.unknown")
	}

	config.initDb()

	config.Port = environment.ParseInt("PORT", config.Port)

	config.D16bXsdPath = environment.Getenv("D16B_XSD_PATH", config.D16bXsdPath)
	config.Ubl21XsdPath = environment.Getenv("UBL21_XSD_PATH", config.Ubl21XsdPath)

	config.SlowStorageType = environment.Getenv("SLOW_STORAGE_TYPE", config.SlowStorageType)
	switch config.SlowStorageType {
	case "local":
		config.LocalStorageBasePath = environment.RequireVar("LOCAL_STORAGE_BASE_PATH")
	case "gcs":
		config.GcsBucket = environment.RequireVar("GCS_BUCKET")
	default:
		log.WithField("slow_storage_type", config.SlowStorageType).Fatal("config.slow_storage_type.unknown")
	}

	log.Info("config.loaded")

	return config
}
