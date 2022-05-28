package serve

import (
	"encoding/json"
	"fmt"
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/vicmanbrile/moneyGolang/db"
	"github.com/vicmanbrile/moneyGolang/serve/auth"
	"github.com/vicmanbrile/moneyGolang/serve/schemas"
)

type AllCredits struct {
	NameProfile string            `json:"profile"`
	Credits     []schemas.Resumen `json:"credits"`
	MoneyInDays float64           `json:"money"`
}

type ErrorNotFound struct {
	Type  int   `json:"type"`
	Error error `json:"error"`
}

/*
	ShowCredits() es un Handler que responde con un AllCredis
*/

func ShowCredits(w http.ResponseWriter, r *http.Request) {

	extractData, err := db.GetDataProfile("6215c7dc38821f527b019d3e", "profile") // Extraemos con un Id y la Collecction de un Perfil
	if err != nil {
		w.WriteHeader(http.StatusNotFound)

		Error := ErrorNotFound{
			Type:  http.StatusNotFound,
			Error: err,
		}

		json.NewEncoder(w).Encode(Error)
	}

	data := AllCredits{
		NameProfile: "vicmanbrile",
		Credits:     extractData.Wallets.Expenses.CalcPerfil(extractData.Wallets.Average),
		MoneyInDays: extractData.Registers.Budgets(),
	}

	{
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(data)
	}

}

func GoServer() {
	PORT := ":8080"

	ApiSubdomain := "api.localhost"
	AssetsSubdomain := "assets.localhost"
	Auth0Subdomain := "auth.localhost"

	{ // API Rest
		http.HandleFunc(ApiSubdomain+"/", ShowCredits)
	}

	{ // Auth0
		http.HandleFunc(Auth0Subdomain+"/public", auth.AllAcces)

		http.Handle(Auth0Subdomain+"/privade", auth.EnsureValidTolken()(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// CORS Headers.
				w.Header().Set("Access-Control-Allow-Credentials", "true")
				w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
				w.Header().Set("Access-Control-Allow-Headers", "Authorization")

				w.Header().Set("Content-Type", "application/json")

				tocken := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)

				claims := tocken.CustomClaims.(*auth.CustomClaims)

				if !claims.HasScope("read:messages") {
					w.WriteHeader(http.StatusForbidden)
					w.Write([]byte(`{"message":"Alcanze no logrado."}`))
					return
				}
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"message":"Hola a un punto publico! Necesitas acceder para ver esto."}`))
			}),
		))
	}

	{ // Assets Request
		http.Handle(AssetsSubdomain+"/", http.StripPrefix("/", http.FileServer(http.Dir("./assets"))))
	}

	fmt.Println("Server listing... http:localhost" + PORT)
	http.ListenAndServe(PORT, nil)
}
