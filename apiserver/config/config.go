package config

import (
	"github.com/slovak-egov/einvoice/environment"
	"log"
	"os"
)

type dbConfiguration struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	InstanceConnectionName string
}

type configuration struct {
	Db                   dbConfiguration
	Port                 int
	D16bXsdPath          string
	Ubl21XsdPath         string
	SlowStorageType      string
	LocalStorageBasePath string
	GcsBucket            string
}

var Config = configuration{}

func InitConfig() {
	Config.Db = dbConfiguration{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     environment.RequireVar("DB_NAME"),
		User:     environment.RequireVar("DB_USER"),
		Password: environment.RequireVar("DB_PASSWORD"),
		InstanceConnectionName: os.Getenv("DB_INSTANCE_CONNECTION_NAME"),
	}

	Config.Port = environment.ParseInt("PORT")

	Config.D16bXsdPath = environment.RequireVar("D16B_XSD_PATH")
	Config.Ubl21XsdPath = environment.RequireVar("UBL21_XSD_PATH")

	Config.SlowStorageType = os.Getenv("SLOW_STORAGE_TYPE")
	switch Config.SlowStorageType {
	case "local":
		Config.LocalStorageBasePath = environment.RequireVar("LOCAL_STORAGE_BASE_PATH")
	case "gcs":
		Config.GcsBucket = environment.RequireVar("GCS_BUCKET")
	default:
		log.Fatal("Unknown SLOW_STORAGE_TYPE:", Config.SlowStorageType)
	}

	log.Println("Config loaded")
}
