package config

import (
	"github.com/sirupsen/logrus"
	"time"
)

var devConfig = configuration{
	Port:         8082,
	RedisUrl:     "localhost:6379",
	ApiServerUrl: "http://localhost:8081",
	SlovenskoSk: SlovenskoSkConfiguration{
		Url: "https://upvs.dev.filipsladek.com",
	},
	LogLevel:        logrus.DebugLevel,
	TokenExpiration: 1 * time.Hour,
}

var prodConfig = configuration{
	SlovenskoSk: SlovenskoSkConfiguration{
		Url: "https://upvs.dev.filipsladek.com",
	},
	LogLevel:        logrus.InfoLevel,
	TokenExpiration: 1 * time.Hour,
}
