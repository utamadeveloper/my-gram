package main

import (
	"fmt"
	"log"
	"my-gram/config"
	"my-gram/router"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.DatabaseConnect()

	app := fmt.Sprintf(
		"%s:%s",
		os.Getenv("APP_HOST"),
		os.Getenv("APP_PORT"),
	)

	r := router.StartApp()
	r.Run(app)
}
