package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func init() {

	err := godotenv.Load() // The Original .env

	if err != nil {
		log.Fatal("Error loading .env file")
		panic(err.Error())
	}

	env := os.Getenv("APP_ENV")
	if len(env) < 1 {
		env = "dev"
	}

	godotenv.Load(".env." + env)
	if err != nil {
		log.Fatal("Error loading .env." + env + " file")
		panic(err.Error())
	}

	if os.Getenv("APP_ENV") != "prod" {
		// Show application environment variables
		for _, e := range os.Environ() {
			pair := strings.SplitN(e, "=", 2)
			if strings.HasPrefix(pair[0], "APP_") {
				fmt.Printf("%s: %s\n", pair[0], pair[1])
			}
		}
	}
}

func main() {
	a := App{}
	a.Initialize()
	a.Run(":" + os.Getenv("APP_HOST_PORT"))
}
