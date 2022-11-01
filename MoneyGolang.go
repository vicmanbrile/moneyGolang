package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type MoneyGolang struct {
	Router *chi.Mux
}

func (mg *MoneyGolang) ListenAndServe() {
	mg.Router = chi.NewRouter()

	mg.Router.Get("/", func(w http.ResponseWriter, r *http.Request) {

		fm := &struct {
			name     string
			lasrname string
		}{
			name:     "Victor",
			lasrname: "Briseño",
		}

		fmt.Println(fm)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(fm); err != nil {
			render.Render(w, r, NewInternalServerError(err))
		}

	})

	/*
		mg.Router.Use(middleware.Logger)

		mg.Router.Get("/", handlers.ShowCredits)

		handlers.FileServer(mg.Router)

		mg.Router.Post("/user", handlers.SessionForm)
		r.Get("/user", handlers.SessionFormGet)

	*/

	log.Fatal(http.ListenAndServe(":8080", mg.Router))
}

func NewInternalServerError(s error) render.Renderer {

	sad := render.Renderer{s}

	return sad

}
