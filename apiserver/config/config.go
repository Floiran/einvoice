package config

import (
	"log"
	"os"

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
}

func (config *Configuration) initDb() {
	config.Db = DbConfiguration{
		Host:     environment.RequireVar("DB_HOST"),
		Port:     environment.ParseInt("DB_PORT"),
		Name:     environment.RequireVar("DB_NAME"),
		User:     environment.RequireVar("DB_USER"),
		Password: environment.RequireVar("DB_PASSWORD"),
	}
}

func Init() Configuration {
	config := Configuration{}
	config.initDb()

	config.Port = environment.ParseInt("PORT")

	config.D16bXsdPath = environment.RequireVar("D16B_XSD_PATH")
	config.Ubl21XsdPath = environment.RequireVar("UBL21_XSD_PATH")

	config.SlowStorageType = os.Getenv("SLOW_STORAGE_TYPE")
	switch config.SlowStorageType {
	case "local":
		config.LocalStorageBasePath = environment.RequireVar("LOCAL_STORAGE_BASE_PATH")
	case "gcs":
		config.GcsBucket = environment.RequireVar("GCS_BUCKET")
	default:
		log.Fatal("Unknown SLOW_STORAGE_TYPE:", config.SlowStorageType)
	}

	log.Println("Config loaded")
	return config
}
