package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	MOUNTH float64 = 28
)

type Perfil struct {
	Creditos     []Product     `json:"credit"`
	Deudas       []Debt        `json:"debts"`
	Suscriptions []Suscription `json:"suscriptions"`
}

type Report interface {
	PriceMount() float64
}

func main() {

	file, err := ioutil.ReadFile("filename.json")
	if err != nil {
		fmt.Printf("Error al convertir a JSON: %v", err)
	}

	data := &Perfil{}

	err = json.Unmarshal([]byte(file), &data)
	if err != nil {
		fmt.Printf("Error al convertir a JSON: %v", err)
	}

	for _, value := range data.Creditos {
		MonthlyPayment(value)
		MonthlyDaysPayment(value)
	}

	for _, value := range data.Deudas {
		MonthlyPayment(value)
		MonthlyDaysPayment(value)
	}

	for _, value := range data.Suscriptions {
		MonthlyPayment(value)
		MonthlyDaysPayment(value)
	}

}

func MonthlyPayment(r Report) {
	total := r.PriceMount()

	fmt.Printf("Pagar $%.2f en un mes.\n", total)
}

func MonthlyDaysPayment(r Report) {
	total := r.PriceMount() / MOUNTH
	fmt.Printf("Pagar $%.2f por día.\n", total)
}
