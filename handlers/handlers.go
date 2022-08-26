package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/vicmanbrile/moneyGolang/db"
	"github.com/vicmanbrile/moneyGolang/schemas"
)

type ErrorNotFound struct {
	Type  int   `json:"type"`
	Error error `json:"error"`
}

// Fileserver for assents folder

func FileServer(r chi.Router) {
	path, root := "/assets", http.Dir("assets")

	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}

/*
	ShowCredits() es un Handler que responde con un AllCredis
*/

type AllCredits struct {
	NameProfile  string
	Credits      []schemas.Resumen
	MoneyInDays  float64
	StyleRecurse template.URL
}

func ShowCredits(w http.ResponseWriter, r *http.Request) {

	chi.URLParam(r, "")

	Cookie, _ := r.Cookie("Profile")

	User := db.Profile{
		ID: Cookie.Value,
	}

	{
		extractData, err := db.GetDataProfile(User) // Extraemos con un Id y la Collecction de un Perfil
		if err != nil {
			w.WriteHeader(http.StatusNotFound)

			Error := ErrorNotFound{
				Type:  http.StatusNotFound,
				Error: err,
			}

			json.NewEncoder(w).Encode(Error)
		}

		data := AllCredits{
			NameProfile:  "vicmanbrile",
			Credits:      extractData.Wallets.Expenses.CalcPerfil(extractData.Wallets.Average),
			MoneyInDays:  extractData.Registers.Budgets(),
			StyleRecurse: template.URL("http://localhost:8080/assets/main.css"),
		}

		Home, _ := template.ParseFiles("templates/show-credits.gohtml")

		Home.Execute(w, data)

	}

}

func SessionForm(w http.ResponseWriter, r *http.Request) {

	user := r.FormValue("user")
	pass := r.FormValue("pass")

	if user == "vicmanbrile" && pass == "Fenian_135" {
		http.SetCookie(w, &http.Cookie{
			Name:     "Profile",
			Value:    "6215c7dc38821f527b019d3e",
			HttpOnly: true,
			Expires:  time.Date(2022, 9, 0, 0, 0, 0, 0, time.UTC),
		})
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func SessionFormGet(w http.ResponseWriter, r *http.Request) {

	FormTemplate, err := template.ParseFiles("templates/form-session.gohtml")

	if err != nil {
		fmt.Println(err)
	}

	_ = FormTemplate.Execute(w, nil)

}

// func AddCredit(w http.ResponseWriter, r *http.Request) {}
