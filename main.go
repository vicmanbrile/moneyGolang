package main

import (
	"encoding/json"
	"flag"
	"fmt"

	database_mongodb "github.com/vicmanbrile/moneyGolang/database"
	"github.com/vicmanbrile/moneyGolang/profile"
	"github.com/vicmanbrile/moneyGolang/status"
)

func main() {

	document, err := database_mongodb.Connection()

	flag.Parse()
	Registro := flag.Args()

	data := &profile.Perfil{}
	err = json.Unmarshal(document, &data)
	if err != nil {
		fmt.Printf("Error al convertir a JSON: %v", err)
	}

	data.PrintTable()

	status.Resumen(Registro...)
}
