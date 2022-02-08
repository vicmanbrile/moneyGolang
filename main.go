package main

import (
	"encoding/json"
	"fmt"

	"github.com/vicmanbrile/moneyGolang/db"
	"github.com/vicmanbrile/moneyGolang/profile"
)

func main() {
	app := Init()

	D := app.Expenses.CalcPerfil(app.Wallets.Average)
	D.PrintTable()

	fmt.Println(app.Free())

}

func Init() (d *profile.Perfil) {
	err := json.Unmarshal(db.GetData(), &d)

	if err != nil {
		fmt.Printf("Error al convertir a JSON: %v", err)
	}

	return
}
