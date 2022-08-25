package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/vicmanbrile/moneyGolang/handlers"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading the .env file: %v", err)
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", handlers.Home)

	r.Get("/credits", handlers.ShowCredits)

	r.Post("/user", handlers.SessionForm)
	r.Get("/user", handlers.SessionFormGet)

	http.ListenAndServe(":8080", r)

	/*

		SER := serve.NuevoServidor(":8080")

		SER.Handle("/credits", handlers.ShowCredits, "POST", middlewares.Logging())

		SER.GoServer()

	*/

}
