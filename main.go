package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	DAYS_MOUNTH  float64 = 28
	MOUNTHS_YEAR float64 = 12
)

type Perfil struct {
	Creditos     []Product     `json:"credit"`
	Deudas       []Debt        `json:"debts"`
	Suscriptions []Suscription `json:"suscriptions"`
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

	fmt.Printf("$%.2f\n", data.PriceMount())

}

func (p *Perfil) PriceMount() float64 {
	var total float64

	for _, value := range p.Creditos {
		total += value.PriceMount()
	}
	for _, value := range p.Deudas {
		total += value.PriceMount()
	}
	for _, value := range p.Suscriptions {
		total += value.PriceMount()
	}

	return total
}
