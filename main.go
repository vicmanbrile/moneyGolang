package main

import (
	"encoding/json"
	"fmt"
	"math"

	"github.com/vicmanbrile/moneyGolang/db"
	"github.com/vicmanbrile/moneyGolang/profile"
)

func main() {
	app := Init()

	matNegative := math.Abs(-5)
	matPositive := math.Abs(5)

	fmt.Println(matNegative, matPositive)

	app.PrintTable()
	app.StutusTable()
}

func Init() (d *profile.Perfil) {
	err := json.Unmarshal(db.GetData(), &d)

	if err != nil {
		fmt.Printf("Error al convertir a JSON: %v", err)
	}

	return
}
