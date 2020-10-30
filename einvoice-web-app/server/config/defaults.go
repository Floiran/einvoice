package config

import "github.com/sirupsen/logrus"

var devConfig = Configuration{
	Port: 8080,
	AuthServerUrl: "http://localhost:8082",
	ClientBuildDir: "einvoice-web-app/client/build",
	SlovenskoSkLoginUrl: "https://upvs.dev.filipsladek.com/login?callback=https://web-app.dev.filipsladek.com/login-callback",
	LogLevel: logrus.DebugLevel,
}

var prodConfig = Configuration{
	ClientBuildDir: "../client/build/",
	SlovenskoSkLoginUrl: "https://upvs.dev.filipsladek.com/login?callback=https://web-app.dev.filipsladek.com/login-callback",
	LogLevel: logrus.InfoLevel,
}
