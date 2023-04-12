package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	router "github.com/vicmanbrile/moneyGolang/api"
	"github.com/vicmanbrile/moneyGolang/api/db"
)

var (
	Port     string = "8080"
	DATABASE *db.MongoConnection
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading the .env file: %v", err)
	}

	r, db := router.NewMoneyRouter()

	defer db.CancelConection()

	fmt.Printf("Server in http://localhost:%v\n", Port)
	log.Fatal(http.ListenAndServe(":"+Port, r))

}
