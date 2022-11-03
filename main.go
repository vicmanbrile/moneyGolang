package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/vicmanbrile/moneyGolang/db"
)

var (
	Port string = "8000"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading the .env file: %v", err)
	}

	money_app := &MoneyGolang{}

	mongo_connection := db.NewMongoConnection()

	defer mongo_connection.CancelConection()

	money_app.ClientDB = mongo_connection

	money_app.ListenAndServe(Port)

}
