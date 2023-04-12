package router

import (
	"encoding/json"
	"fmt"
	"github.com/vicmanbrile/moneyGolang/api/middlewares"
	"html/template"
	"io"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/vicmanbrile/moneyGolang/api/db"
	"github.com/vicmanbrile/moneyGolang/api/schemas"
	"github.com/vicmanbrile/moneyGolang/application"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ErrorNotFound struct {
	Type  int   `json:"type"`
	Error error `json:"error"`
}

// File-server for assents folder

func FileServer(r chi.Router) {
	path, root := "/assets", http.Dir("web")

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
	ShowCredits() es un Handler que responde con un AllCredits
*/

type AllCredits struct {
	NameProfile  string
	Credits      []schemas.Resumen
	MoneyInDays  float64
	StyleRecurse template.URL
	Success      bool
}

func ApiRouter(r chi.Router) {
	r.Get("/person", middlewares.Loggin(func(w http.ResponseWriter, r *http.Request) {

		// Struct to return
		var person struct {
			Name         string            `json:"name"`
			Credits      []schemas.Resumen `json:"credits"`
			MoneyInDays  float64           `json:"money_in_days"`
			StyleRecurse template.URL      `json:"style-recurse"`
			Success      bool              `json:"success"`
		}

		// Search for User and Password (Change to middleware login and add to result to cookie)
		objectId, err := db.SeachUser("vicmanbrile", "Fenian_135")
		if err != nil {
			fmt.Println(err)
		}

		find := &db.User{
			ID: objectId,
		}

		extractData := find.ReadProfile(database) // Extraemos id y la Collection de un Perfil
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

		person.Name = "vicmanbrile"
		person.Credits = extractData.Wallets.Expenses.CalcPerfil(extractData.Wallets.Average)
		person.MoneyInDays = extractData.Budgets()
		person.StyleRecurse = template.URL("http://localhost:8080/assets/main.css")
		person.Success = false

		w.Header().Set("Content-Type", "application/json")

		personJson, _ := json.Marshal(person)
		w.Write(personJson)
	}))

	r.Post("/person", func(writer http.ResponseWriter, request *http.Request) {
		var Document struct {
			User string `json:"user"`
		}

		d, _ := io.ReadAll(request.Body)

		defer request.Body.Close()

		json.Unmarshal(d, &Document)

		writer.Header().Add("Content-Type", "application/json")

		s, _ := json.Marshal(Document)

		fmt.Fprintf(writer, "%s", s)
	})
}

func ShowCredits() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		err := r.ParseForm()
		if err != nil {
			return
		}

		// Cookie, _ := r.Cookie("Profile")

		objectId, err := db.SeachUser("vicmanbrile", "Fenian_135")
		if err != nil {
			fmt.Println(err)
		}

		find := &db.User{
			ID: objectId,
		}

		{
			extractData := find.ReadProfile(database) // Extraemos id y la Collection de un Perfil
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

			var (
				data = AllCredits{
					NameProfile:  "vicmanbrile",
					Credits:      extractData.Wallets.Expenses.CalcPerfil(extractData.Wallets.Average),
					MoneyInDays:  extractData.Budgets(),
					StyleRecurse: template.URL("http://localhost:8080/assets/main.css"),
					Success:      false,
				}
			)
			var files = []string{
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
		w.Header().Add("Authorization", "Basic dmljbWFuYnJpbGU6RmVuaWFuXzEzNQ==")
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

func AveragePost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		insert := &db.User{}

		ds := application.Deposits{
			YearDay:  200,
			Deposits: 18000,
		}

		insert.InsertDeposit(database, ds)

		fmt.Fprintf(w, "Hello World!")
	}
}

func AverageGet() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		id, err := primitive.ObjectIDFromHex("6362b84b70a43aee546d8745")
		if err != nil {
			fmt.Println(err)
		}

		find := &db.User{
			ID: id,
		}

		depost := find.ReadDeposit(database)

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
