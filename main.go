package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/vicmanbrile/moneyGolang/serve"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading the .env file: %v", err)
	}

	serve.GoServer()
}
