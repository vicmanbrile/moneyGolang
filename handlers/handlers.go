package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/vicmanbrile/moneyGolang/application"
	"github.com/vicmanbrile/moneyGolang/db"
	"github.com/vicmanbrile/moneyGolang/handlers/schemas"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		r.Get(path, http.RedirectHandler(path+"/", http.StatusMovedPermanently).ServeHTTP)
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
	Success      bool
}

func ShowCredits(ClientDB *db.MongoConnection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		err := r.ParseForm()
		if err != nil {
			return
		}

		// Cookie, _ := r.Cookie("Profile")

		objectId, err := primitive.ObjectIDFromHex("6215c7dc38821f527b019d3e")
		if err != nil {
			fmt.Println(err)
		}

		find := &db.User{
			ID: objectId,
		}

		{
			extractData := find.ReadProfile(ClientDB) // Extraemos con un id y la Collecction de un Perfil
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				Error := ErrorNotFound{
					Type:  http.StatusNotFound,
					Error: err,
				}

				err := json.NewEncoder(w).Encode(Error)
				if err != nil {
					return
				}
			}

			fmt.Println(extractData)

			data := AllCredits{
				NameProfile:  "vicmanbrile",
				Credits:      extractData.Wallets.Expenses.CalcPerfil(extractData.Wallets.Average),
				MoneyInDays:  extractData.Budgets(),
				StyleRecurse: template.URL("http://localhost:8080/assets/main.css"),
				Success:      false,
			}

			files := []string{
				"./templates/main.gohtml",
				"./templates/show-credits.gohtml",
				"./templates/componets/navegation-bar.gohtml",
			}

			Home, _ := template.ParseFiles(files...)

			Home.Execute(w, data)

		}
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

func SessionFormGet(w http.ResponseWriter, _ *http.Request) {

	files := []string{
		"./templates/main.gohtml",
		"./templates/form-session.gohtml",
		"./templates/componets/navegation-bar.gohtml",
	}

	FormTemplate, err := template.ParseFiles(files...)

	if err != nil {
		fmt.Println(err)
	}

	_ = FormTemplate.Execute(w, nil)

}

func AveragePost(ClientDB *db.MongoConnection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		insert := &db.User{}

		ds := application.Deposits{
			YearDay:  200,
			Deposits: 18000,
		}

		insert.InsertDeposit(ClientDB, ds)

		fmt.Fprintf(w, "Hello World!")
	}
}

func AverageGet(ClientDB *db.MongoConnection) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		id, err := primitive.ObjectIDFromHex("6362b84b70a43aee546d8745")
		if err != nil {
			fmt.Println(err)
		}

		find := &db.User{
			ID: id,
		}

		depost := find.ReadDeposit(ClientDB)

		rest := &schemas.MostrarDeposits{
			Average:  depost.Average(),
			YearDay:  depost.YearDay,
			Deposits: depost.Deposits,
		}

		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusCreated)

		err = json.NewEncoder(w).Encode(rest)
		if err != nil {
			return
		}
	}

}
