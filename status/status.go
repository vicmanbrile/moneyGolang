package status

import (
	"os"
	"time"

	"github.com/olekukonko/tablewriter"
)

type Registers struct {
	Spent []struct {
		Key   string  `json:"key"`
		Value float64 `json:"value"`
	} `json:"spent"`
	Saved []struct {
		Key   string  `json:"key"`
		Value float64 `json:"value"`
	} `json:"saved"`
	Extras []struct {
		Week  float64 `json:"week"`
		Extra float64 `json:"extra"`
		Days  float64 `json:"days"`
	} `json:"extras"`
}

type Budget struct {
	Total    float64
	Restados float64
	Dias     float64
}

func (r *Registers) BudgetsNow(money float64) (Bdgt Budget) {
	Bdgt.Dias = automaticTime()
	Bdgt.Total = Bdgt.Dias * money

	{
		var restar float64
		for _, value := range r.Saved {
			restar += value.Value
		}

		for _, value := range r.Saved {
			restar += value.Value
		}

		Bdgt.Restados = restar
	}

	return
}

func (r *Registers) BudgetsWon(money float64) (Bdgt Budget) {
	var days float64
	for _, value := range r.Extras {
		days += value.Days
	}
	{
		var restar float64
		for _, value := range r.Saved {
			restar += value.Value
		}

		for _, value := range r.Saved {
			restar += value.Value
		}

		Bdgt.Restados = restar
	}

	Bdgt.Dias = days
	Bdgt.Total = days * money
	return
}

func (r *Registers) Resumen(saldo float64) [][]string {
	info := make([][]string, 0)

	{
		var row []string
		dgts := r.BudgetsNow(saldo)
		row = append("", "", "", "")
	}

	{
		r.BudgetsWon(saldo)
	}

	return info
}

func (r *Registers) PrintTable(money float64) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetBorder(false)

	table.SetHeader([]string{"Dias", "Debemos", "Falta", "Estatus"})

	info := r.Resumen(money)

	table.SetFooter([]string{
		"",
		"",
		"",
		"",
	})

	for _, v := range info {
		table.Append(v)
	}

	table.Render()

}

func (b *Budget) Lack(percentage float64) float64 {
	return (b.Total * percentage) - b.Restados
}

func (b *Budget) Free(percentage float64) float64 {
	return b.Total * percentage
}

func automaticTime() float64 {
	today := time.Now().YearDay()

	return float64(today)
}
