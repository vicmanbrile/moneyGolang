package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/vicmanbrile/moneyGolang/db"
	"github.com/vicmanbrile/moneyGolang/handlers"
	"github.com/vicmanbrile/moneyGolang/handlers/middlewares"
	"github.com/vicmanbrile/moneyGolang/handlers/middlewares/auth"
)

type MoneyGolang struct {
	Router   *chi.Mux
	ClientDB *db.MongoConnection
}

func (mg *MoneyGolang) ListenAndServe(port string) {
	mg.Router = chi.NewRouter()

	mg.Router.Use(middleware.Logger, middlewares.Loggin)
	mg.Router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	mg.Router.Get("/", handlers.Test)
	mg.Router.Get("/api/public", handlers.AllAcces)
	mg.Router.Handle("/api/private", auth.EnsureValidToken()(http.HandlerFunc(handlers.PrivateAccess)))
	mg.Router.Handle("/api/private-scoped", auth.EnsureValidToken()(http.HandlerFunc(handlers.PrivateAccessScoped)))

	/*
		 mg.Router.Get("/", handlers.ShowCredits(mg.ClientDB))

		handlers.FileServer(mg.Router)

		mg.Router.Post("/average", handlers.AveragePost(mg.ClientDB))
		mg.Router.Get("/average", handlers.AverageGet(mg.ClientDB))

		mg.Router.Post("/user", handlers.SessionForm)
		mg.Router.Get("/user", handlers.SessionFormGet)
	*/

	log.Fatal(http.ListenAndServe(":"+port, mg.Router))
}

/*
func registerAPI(r *chi.Mux) {
	s := oauth.NewBearerServer(
		"mySecretKey-10101",
		time.Second*120,
		nil,
		nil)

	r.Post("/token", s.UserCredentials)
	r.Post("/auth", s.ClientCredentials)
}

*/
