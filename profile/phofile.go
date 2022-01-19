package profile

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/vicmanbrile/moneyGolang/status"
)

var (
	DAYS_MOUNTH  float64 = 30
	MOUNTHS_YEAR float64 = 12
)

type Perfil struct {
	Creditos     []Credit         `json:"credit"`
	Deudas       []Debt           `json:"debts"`
	Suscriptions []Suscription    `json:"suscriptions"`
	Wallets      Wallet           `json:"wallets"`
	Percentiles  []Percentile     `json:"percentile"`
	Registers    status.Registers `json:"registers"`
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

	for _, value := range p.Percentiles {
		total += value.PriceMount() * p.Wallets.Average
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

	for _, value := range p.Percentiles {
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
