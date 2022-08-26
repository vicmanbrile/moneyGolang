package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/vicmanbrile/moneyGolang/handlers"
	"log"
	"net/http"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading the .env file: %v", err)
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", handlers.ShowCredits)

	handlers.FileServer(r)

	r.Post("/user", handlers.SessionForm)
	r.Get("/user", handlers.SessionFormGet)

	log.Fatal(http.ListenAndServe(":8080", r))

}
