package expenses

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

var (
	DAYS_MOUNTH  float64 = 30
	MOUNTHS_YEAR float64 = 12
)

type Expenses struct {
	Creditos     []Credit      `json:"credit"`
	Deudas       []Debt        `json:"debts"`
	Suscriptions []Suscription `json:"suscriptions"`
	Percentiles  []Percentile  `json:"percentile"`
}

func (e *Expenses) CalcPerfil(Average float64) []Resumen {
	var Todos []Resumen

	for _, value := range e.Creditos {
		Todos = append(Todos, *value.CalcCredit(Average))
	}
	for _, value := range e.Deudas {
		Todos = append(Todos, *value.CalcDebt(Average))
	}

	for _, value := range e.Suscriptions {
		Todos = append(Todos, *value.CalcSuscriptions(Average))
	}

	for _, value := range e.Percentiles {
		Todos = append(Todos, *value.CalcPercentiles(Average))
	}

	return Todos
}

func (e *Expenses) PriceDays(Average float64) float64 {
	var result float64

	for _, value := range e.CalcPerfil(Average) {
		result += value.PriceDay()
	}

	return result
}

func (e *Expenses) PrintTable(Average float64) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetBorder(false)

	table.SetHeader([]string{"Grupo", "Descripcion", "Porcentaje", "$ X D", "Complete"})

	info := make([][]string, 0)
	{
		for _, value := range e.CalcPerfil(Average) {
			d := value.Resumen(Average)
			info = append(info, d)
		}
	}

	table.SetFooter([]string{
		"",
		"Total:",
		fmt.Sprintf("%.2f%%", (e.PriceDays(Average)/Average)*100),
		fmt.Sprintf("$%.2f", e.PriceDays(Average)),
		"",
	})

	for _, v := range info {
		table.Append(v)
	}

	table.Render()
}

type Resumen struct {
	PriceYear  float64
	Porcentile float64
	Complete   float64
	Name       string
	Type       string
}

func (r *Resumen) PriceDay() float64 {
	return r.PriceYear / (DAYS_MOUNTH * MOUNTHS_YEAR)
}

func (r *Resumen) Resumen(salary float64) []string {
	info := make([]string, 5)

	info[0] = r.Type
	info[1] = r.Name
	info[2] = fmt.Sprintf("%%%.2f", r.Porcentile*100)
	info[3] = fmt.Sprintf("$%.2f", r.PriceDay())
	info[4] = fmt.Sprintf("%%%.2f", r.Complete*100)

	return info
}
