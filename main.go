package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading the .env file: %v", err)
	}

	//Revisar si existen las colecciones : Deposits, Shoppings, Suscriptions, Wallet

	MG := &MoneyGolang{}

	defer MG.CloseDatabase()

	MG.ListenAndServe()

}
