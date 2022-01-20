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

func (p *Perfil) CalcPerfil() []Resumen {
	var Todos []Resumen

	for _, value := range p.Creditos {
		Todos = append(Todos, *value.CalcCredit())
	}
	for _, value := range p.Deudas {
		Todos = append(Todos, *value.CalcDebt())
	}

	for _, value := range p.Suscriptions {
		Todos = append(Todos, *value.CalcSuscriptions())
	}

	for _, value := range p.Percentiles {
		Todos = append(Todos, *value.CalcPercentiles(p.Wallets.Average))
	}

	return Todos
}

func (p *Perfil) PriceDays() float64 {
	var result float64

	for _, value := range p.CalcPerfil() {
		result += value.PriceForDays()
	}

	return result
}

func (p *Perfil) Resumen() [][]string {
	info := make([][]string, 0)

	for _, value := range p.CalcPerfil() {
		d := value.Resumen(p.Wallets.Average)
		info = append(info, d)
	}

	return info
}

func (p *Perfil) PrintTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetBorder(false)

	table.SetHeader([]string{"Grupo", "Descripcion", "Porcentaje", "$ X D"})

	info := p.Resumen()

	table.SetFooter([]string{
		"",
		"Total:",
		fmt.Sprintf("%.2f%%", (p.PriceDays()/p.Wallets.Average)*100),
		fmt.Sprintf("$%.2f", p.PriceDays()),
	})

	for _, v := range info {
		table.Append(v)
	}

	table.Render()
}

type Resumen struct {
	PriceMount float64
	Name       string
	Type       string
}

func (r *Resumen) PriceForDays() float64 {
	return r.PriceMount / DAYS_MOUNTH
}

func (r *Resumen) Resumen(salary float64) []string {
	info := make([]string, 4)

	info[0] = r.Type
	info[1] = r.Name
	info[2] = fmt.Sprintf("%.2f%%", (r.PriceForDays()/salary)*100)
	info[3] = fmt.Sprintf("$%.2f", r.PriceForDays())

	return info
}
