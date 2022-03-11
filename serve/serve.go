package serve

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/vicmanbrile/moneyGolang/db"
	"github.com/vicmanbrile/moneyGolang/profile/expenses"
)

type AllCredits struct {
	PageTitle string
	Todos     expenses.AllExpenses
}

func ShowCredits(w http.ResponseWriter, r *http.Request) {
	tml := template.Must(template.ParseFiles("serve/public/index.html"))
	extractData := db.Init()

	data := AllCredits{
		PageTitle: "Mis Creditos",
		Todos:     extractData.Wallets.Expenses.CalcPerfil(extractData.Wallets.Average),
	}

	tml.Execute(w, data)
}

func GetData() {
	http.HandleFunc("/", ShowCredits)

	fs := http.FileServer(http.Dir("./assets"))

	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	fmt.Println("Server listing... http:localhost:8080")
	http.ListenAndServe(":8080", nil)
}

type Todo struct {
	Title string
	Done  bool
}
