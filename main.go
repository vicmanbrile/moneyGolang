package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading the .env file: %v", err)
	}

	MG := &MoneyGolang{}

	MG.ListenAndServe()

}
