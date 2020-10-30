package config

import "github.com/sirupsen/logrus"

var devConfig = configuration{
	Port: 8082,
	RedisUrl: "localhost:6379",
	ApiServerUrl: "http://localhost:8081",
	SlovenskoSk: SlovenskoSkConfiguration{
		Url: "https://upvs.dev.filipsladek.com",
	},
	LogLevel: logrus.DebugLevel,
}

var prodConfig = configuration{
	SlovenskoSk: SlovenskoSkConfiguration{
		Url: "https://upvs.dev.filipsladek.com",
	},
	LogLevel: logrus.InfoLevel,
}
