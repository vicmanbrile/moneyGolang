package main

import (
	"fmt"

	database_mongodb "github.com/vicmanbrile/moneyGolang/database"
)

func main() {

	app := database_mongodb.Data{}
	app.Init()

	app.Perfil.PrintTable()
	A := app.Perfil.Registers.BudgetsNow(app.Perfil.Wallets.Average)
	B := app.Perfil.Registers.BudgetsWon(app.Perfil.Wallets.Average)
	{
		fmt.Printf("Dinero para ahora:%.2f , Dinero de entregado: %.2f\n", A.Total, B.Total)
	}
	{
		C := A.Lack((app.Perfil.PriceDays() / app.Perfil.Wallets.Average))
		D := B.Lack((app.Perfil.PriceDays() / app.Perfil.Wallets.Average))
		fmt.Printf("Dinero para ahora tenemos:%.2f , Dinero de entregado tenemos: %.2f\n", C, D)
	}
	{
		C := A.Free(1 - (app.Perfil.PriceDays() / app.Perfil.Wallets.Average))
		D := B.Free(1 - (app.Perfil.PriceDays() / app.Perfil.Wallets.Average))
		fmt.Printf("Dinero para ahora Libres:%.2f , Dinero de entregado Libres: %.2f\n", C, D)
	}

}
