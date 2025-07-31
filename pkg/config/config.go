package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

var (
	appHostKey = "APP_HOST"
	AppHost    string // APP_HOST variable

	appPortKey = "APP_PORT"
	AppPort    uint16 // APP_PORT variable

	appIDKey = "APP_ID"
	AppID    string // APP_ID variable

	appSecretKey = "APP_SECRET"
	AppSecret    string // APP_SECRET variable

	databaseURLKey = "DATABASE_URL"
	DatabaseURL    string // DATABASE_URL variable
)

func init() {
	AppHost = getEnv(appHostKey, "0.0.0.0")
	AppID = mustGetEnv(appIDKey)
	AppSecret = mustGetEnv(appSecretKey)
	DatabaseURL = mustGetEnv(databaseURLKey)

	// AppPort setup
	appPortString := mustGetEnv(appPortKey)
	appPortInt, err := strconv.Atoi(appPortString)
	if err != nil {
		log.Fatalf("%s must be a number\n", appPortKey)
	}
	if appPortInt < 0 || appPortInt > 65535 {
		log.Fatalf("%s must be an integer between 0 and 65535", appPortKey)
	}
	AppPort = uint16(appPortInt)
}

func mustGetEnv(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		msg := fmt.Sprintf("%s environment variable was not set", key)
		log.Fatalln(msg)
		panic(msg)
	}
	return val
}

func getEnv(key, fb string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		log.Printf("Warning %s environment variable was not set! Using fallback %s", key, fb)
		return fb
	}
	return val
}
