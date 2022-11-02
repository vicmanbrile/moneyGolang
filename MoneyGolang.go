package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/vicmanbrile/moneyGolang/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MoneyGolang struct {
	Router   *chi.Mux
	Database *db.MongoConnection
}

type MostrarDeposits struct {
	Average float64 `json:"average"`
}

func (mg *MoneyGolang) ListenAndServe() {
	mg.Router = chi.NewRouter()

	mg.Database = db.EstablishingConnection()

	mg.Router.Get("/average", func(w http.ResponseWriter, r *http.Request) {

		id, _ := primitive.ObjectIDFromHex("6215c7dc38821f527b019d3e")

		find := &db.User{
			ID: id,
		}

		depost := find.ReadDeposit(mg.Database)

		rest := &MostrarDeposits{
			Average: depost.Average(),
		}

		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(rest)

	})

	log.Fatal(http.ListenAndServe(":8000", mg.Router))
}

func (mg *MoneyGolang) CloseDatabase() {
	mg.Database.CancelConection()
}

/*
	mg.Router.Use(middleware.Logger)

	mg.Router.Get("/", handlers.ShowCredits)

	handlers.FileServer(mg.Router)

	mg.Router.Post("/user", handlers.SessionForm)
	r.Get("/user", handlers.SessionFormGet)

*/
