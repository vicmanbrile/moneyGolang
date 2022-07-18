package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/vicmanbrile/moneyGolang/db"
	"github.com/vicmanbrile/moneyGolang/schemas"
	"github.com/vicmanbrile/moneyGolang/templates"
)

type ErrorNotFound struct {
	Type  int   `json:"type"`
	Error error `json:"error"`
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

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World\n %s", r.Host)
}

func ShowCredits(w http.ResponseWriter, r *http.Request) {

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
			StyleRecurse: template.URL("http://assets.localhost:8080/main.css"),
		}

		Home := template.New("Home")

		Home.Parse(templates.ShowCredits)

		Home.Execute(w, data)

		/*{
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.WriteHeader(http.StatusAccepted)
			json.NewEncoder(w).Encode(data)
		}*/

	}

}

func SessionForm(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		user := r.FormValue("user")
		pass := r.FormValue("pass")

		if user == "vicmanbrile" && pass == "Fenian_135" {
			http.SetCookie(w, &http.Cookie{
				Name:     "Profile",
				Value:    "6215c7dc38821f527b019d3e",
				HttpOnly: true,
			})
		}

	}

	FormTemplate := template.New("Session Form")
	FormTemplate.Parse(templates.FormSession)

	FormTemplate.Execute(w, nil)
}

// func AddCredit(w http.ResponseWriter, r *http.Request) {}
