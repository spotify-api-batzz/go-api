package main

import (
	"log"

	"github.com/batzz-00/goutils/logger"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database := Database{}
	err = database.Connect()
	if err != nil {
		logger.Log("Failed to connect to database", logger.Error)
		log.Fatal(err)
	}

	server := Server{Database: &database}
	server.Serve()
}
