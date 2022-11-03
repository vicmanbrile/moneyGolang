package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/vicmanbrile/moneyGolang/db"
	"github.com/vicmanbrile/moneyGolang/handlers"
)

type MoneyGolang struct {
	Router   *chi.Mux
	ClientDB *db.MongoConnection
}

func (mg *MoneyGolang) ListenAndServe(port string) {
	mg.Router = chi.NewRouter()

	mg.Router.Use(middleware.Logger)

	mg.Router.Get("/", handlers.ShowCredits(mg.ClientDB))

	handlers.FileServer(mg.Router)

	mg.Router.Post("/average", handlers.AveragePost(mg.ClientDB))
	mg.Router.Get("/average", handlers.AverageGet(mg.ClientDB))

	mg.Router.Post("/user", handlers.SessionForm)
	mg.Router.Get("/user", handlers.SessionFormGet)

	log.Fatal(http.ListenAndServe(":"+port, mg.Router))
}
