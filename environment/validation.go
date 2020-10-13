package environment

import (
	"log"
	"os"
	"strconv"
)

func ParseInt(varName string) int {
	parsedVar, parseError := strconv.Atoi(os.Getenv(varName))
	if parseError != nil {
		log.Fatal(varName, " has to be int")
	}

	return parsedVar
}

func RequireVar(varName string) string {
	envVar, ok := os.LookupEnv(varName)
	if !ok {
		log.Fatal(varName, " has to be defined")
	}

	return envVar
}
