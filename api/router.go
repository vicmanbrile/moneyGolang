package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/vicmanbrile/moneyGolang/api/db"
	"github.com/vicmanbrile/moneyGolang/api/middlewares"
)

var database *db.MongoConnection

func NewMoneyRouter() (*chi.Mux, *db.MongoConnection) {

	database = db.NewMongoConnection()

	// Router

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Route("/api", ApiRouter)

	r.Get("/", middlewares.Loggin(ShowCredits()))

	FileServer(r)

	r.Post("/average", AveragePost())
	r.Get("/average", AverageGet())

	r.Post("/auth", SessionForm)
	r.Get("/auth", SessionFormGet)

	return r, database
}
