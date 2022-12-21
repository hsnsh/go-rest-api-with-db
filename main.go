package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strings"

	. "go-rest-api-with-db/helpers"
)

var appLogger IFileLogger

func init() {

	appLogger = NewFileLogger()

	err := godotenv.Load() // The Original .env
	if err != nil {
		appLogger.Fatal("Error loading .env file")
	}

	env := os.Getenv("APP_ENV")
	if len(env) < 1 {
		env = "dev"
	}

	godotenv.Load(".env." + env)
	if err != nil {
		appLogger.Fatal("Error loading .env." + env + " file")
	}

	if os.Getenv("APP_ENV") != "prod" {
		// Show application environment variables
		for _, e := range os.Environ() {
			pair := strings.SplitN(e, "=", 2)
			if strings.HasPrefix(pair[0], "APP_") {
				appLogger.Info(fmt.Sprintf("%s: %s", pair[0], pair[1]))
			}
		}
	}
}

func main() {
	a := App{}
	a.Initialize(os.Getenv("APP_DB_CONNECTION"))
	a.Run(fmt.Sprintf("%s:%s", os.Getenv("APP_HOST_ADDRESS"), os.Getenv("APP_HOST_PORT")))
}
