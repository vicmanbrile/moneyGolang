package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/olekukonko/tablewriter"
)

var (
	DAYS_MOUNTH  float64 = 28
	MOUNTHS_YEAR float64 = 12
)

type Perfil struct {
	Creditos     []Product     `json:"credit"`
	Deudas       []Debt        `json:"debts"`
	Suscriptions []Suscription `json:"suscriptions"`
	Wallets      Wallet        `json:""wallets`
}

type Wallet struct {
	Cash    float64 `json:"cash"`
	Average float64 `json:"average"`
	Banking float64 `json:"banking"`
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

	data.PrintTable()

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

func (p *Perfil) PriceForDays() float64 {
	return p.PriceMount() / DAYS_MOUNTH
}

func (p *Perfil) Resumen() [][]string {
	info := make([][]string, 0)

	for _, value := range p.Creditos {
		info = append(info, value.Resumen(p.Wallets.Average))
	}

	for _, value := range p.Deudas {
		info = append(info, value.Resumen(p.Wallets.Average))
	}

	for _, value := range p.Suscriptions {
		info = append(info, value.Resumen(p.Wallets.Average))
	}

	return info
}

func (p *Perfil) PrintTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetBorder(false)

	table.SetHeader([]string{"Grupo", "Descripcion", "Porcentaje", "$ X D"})

	info := p.Resumen()

	table.SetFooter([]string{"", "Total:", fmt.Sprintf("%.2f%%", (p.PriceForDays()/p.Wallets.Average)*100), fmt.Sprintf("$%.2f", p.PriceForDays())})

	for _, v := range info {
		table.Append(v)
	}

	table.Render()
}
