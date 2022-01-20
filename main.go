package main

import (
	database_mongodb "github.com/vicmanbrile/moneyGolang/database"
)

func main() {

	app := database_mongodb.Data{}
	app.Init()

	app.Perfil.PrintTable()

	app.Perfil.Registers.PrintTable(app.Perfil.Wallets.Average)
}
