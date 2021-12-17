package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Perfil struct {
	Creditos     []Product     `json:"credit"`
	Deudas       []Debt        `json:"debts"`
	Suscriptions []Suscription `json:"suscriptions"`
}

type Report interface {
	PriceMount() float64
	GetName() string
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

}

func MonthlyPayment(r Report) {
	total := r.PriceMount()

	fmt.Printf("Pagar $%.2f para %s en un mes.\n", total, r.GetName())
}

func MonthlyDaysPayment(r Report) {
	total := r.PriceMount() / 30
	fmt.Printf("Pagar $%.2f a %s. por día.\n", total, r.GetName())
}
