package config

import (
	"time"

	"github.com/sirupsen/logrus"
)

var devConfig = Configuration{
	Port:         8082,
	RedisUrl:     "localhost:6379",
	ApiServerUrl: "http://localhost:8081",
	SlovenskoSk: SlovenskoSkConfiguration{
		Url: "https://upvs.dev.filipsladek.com",
	},
	Db: DbConfiguration{
		Host: "localhost",
		Port: 5432,
		Name: "authproxy",
		User: "postgres",
	},
	LogLevel:        logrus.DebugLevel,
	TokenExpiration: 24 * time.Hour,
}

var prodConfig = Configuration{
	SlovenskoSk: SlovenskoSkConfiguration{
		Url: "https://upvs.dev.filipsladek.com",
	},
	Db: DbConfiguration{
		Port: 5432,
		Name: "einvoice",
	},
	LogLevel:        logrus.InfoLevel,
	TokenExpiration: 1 * time.Hour,
}
