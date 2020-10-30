package environment

import (
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func ParseInt(varName string, defaultValue int) int {
	parsedVar, parseError := strconv.Atoi(
		Getenv(varName, strconv.Itoa(defaultValue)),
	)
	if parseError != nil {
		log.WithFields(log.Fields{
			"env": varName,
			"error": parseError,
		}).Fatal("environment.parse_int.error")
	}

	return parsedVar
}

func RequireVar(varName string) string {
	envVar, ok := os.LookupEnv(varName)
	if !ok {
		log.WithField("env", varName).Fatal("environment.require_var.no_value")
	}

	return envVar
}

func Getenv(varName, defaultValue string) string {
	envVar, ok := os.LookupEnv(varName)

	if !ok {
		log.WithFields(log.Fields{
			"env": varName,
			"defaultValue": defaultValue,
		}).Debug("environment.get_env.default_value")
		return defaultValue
	}

	return envVar
}
