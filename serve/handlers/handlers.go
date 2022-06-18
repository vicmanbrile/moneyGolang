package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vicmanbrile/moneyGolang/db"
	"github.com/vicmanbrile/moneyGolang/serve/schemas"
)

type ErrorNotFound struct {
	Type  int   `json:"type"`
	Error error `json:"error"`
}

/*
	ShowCredits() es un Handler que responde con un AllCredis
*/

type AllCredits struct {
	NameProfile string            `json:"profile"`
	Credits     []schemas.Resumen `json:"credits"`
	MoneyInDays float64           `json:"money"`
}

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

func DocHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("\nPrinting all cookies")
	for _, c := range r.Cookies() {
		fmt.Println(c)
	}

	w.WriteHeader(200)
	w.Write([]byte("Doc Get Successful"))
}
