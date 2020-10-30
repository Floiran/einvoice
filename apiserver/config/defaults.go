package config

import "github.com/sirupsen/logrus"

var devConfig = Configuration{
	Db: DbConfiguration{
		Host: "localhost",
		Port: 5432,
		Name: "einvoice",
		User: "postgres",
	},
	Port: 8081,
	D16bXsdPath: "xml/d16b/xsd",
	Ubl21XsdPath: "xml/ubl21/xsd",
	SlowStorageType: "local",
	LogLevel: logrus.DebugLevel,
}

var prodConfig = Configuration{
	Db: DbConfiguration{
		Port: 5432,
		Name: "einvoice",
	},
	D16bXsdPath: "xml/d16b/xsd",
	Ubl21XsdPath: "xml/ubl21/xsd",
	SlowStorageType: "gcs",
	LogLevel: logrus.InfoLevel,
}

var testConfig = Configuration{
	Db: DbConfiguration{
		Host: "localhost",
		Port: 5432,
		Name: "einvoice",
		User: "postgres",
	},
	Port: 8081,
	D16bXsdPath: "xml/d16b/xsd",
	Ubl21XsdPath: "xml/ubl21/xsd",
	SlowStorageType: "local",
	LogLevel: logrus.InfoLevel,
}
