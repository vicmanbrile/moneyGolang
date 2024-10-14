package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/vicmanbrile/moneyGolang/db"
)

var (
	Port string = "8080"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading the .env file: %v", err)
	}

	moneyApp := &MoneyGolang{}

	mongoConnection := db.NewMongoConnection()

	defer mongoConnection.CancelConection()

	moneyApp.ClientDB = mongoConnection

	moneyApp.ListenAndServe(Port)

}
