package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/vicmanbrile/moneyGolang/handlers"
	"github.com/vicmanbrile/moneyGolang/middlewares"
	"github.com/vicmanbrile/moneyGolang/serve"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading the .env file: %v", err)
	}

	SER := serve.NuevoServidor(":8080")

	SER.Handle("/", handlers.Home, "GET")

	SER.Handle("/credits", handlers.ShowCredits, "GET", middlewares.Logging())
	SER.Handle("/credits", handlers.ShowCredits, "POST", middlewares.Logging())

	SER.Handle("/session", handlers.SessionForm, "POST", middlewares.Logging())
	SER.Handle("/session", handlers.SessionForm, "GET", middlewares.Logging())

	SER.GoServer()

}
