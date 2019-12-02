package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {

	log.Println("Hello, World!")

	log.Printf("Unix timestamp:  %d \n", time.Now().Unix())

	err := godotenv.Load()
	if err != nil {

		log.Fatal("Error loading .env file")
	}

	appName := os.Getenv("APP_NAME")
	isDebug := os.Getenv("APP_DEBUG")

	log.Println("APP_NAME: ", appName)
	log.Println("APP_DEBUGL: ", isDebug)

}
