package main

import (
	"flag"

	database_mongodb "github.com/vicmanbrile/moneyGolang/database"
	"github.com/vicmanbrile/moneyGolang/status"
)

func main() {

	app := database_mongodb.Data{}
	app.Init()

	flag.Parse()
	Registro := flag.Args()

	app.Perfil.PrintTable()

	status.Resumen(Registro...)
}
