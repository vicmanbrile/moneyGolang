package expenses

import (
	"fmt"
	"os"
	"time"

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

func (e *Expenses) FaltaMount(Average float64) float64 {
	var result float64

	for _, value := range e.CalcPerfil(Average) {
		result += value.PayMountNow()
	}

	return result
}

func (e *Expenses) FaltaMountPorcentile(Average float64) float64 {
	var result float64

	for _, value := range e.CalcPerfil(Average) {
		result += value.PriceDayNow() / 240
	}

	return result
}

func (e *Expenses) PrintTable(Average float64) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetBorder(false)

	table.SetHeader([]string{"Grupo", "Descripcion", "Porcentaje", "$ X D", "Falta"})

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
		fmt.Sprintf("%.2f%%", e.FaltaMountPorcentile(Average)*100),
		fmt.Sprintf("$%.2f", e.PriceDays(Average)),
		fmt.Sprintf("$%.2f", e.FaltaMount(Average)),
	})

	for _, v := range info {
		table.Append(v)
	}

	table.Render()
}

type Resumen struct {
	PriceYear   float64
	Porcentile  float64
	Complete    float64
	Name        string
	Type        string
	MountInit   float64
	MountsToPay float64
}

func (r *Resumen) PayMountNow() float64 {
	mount := float64(time.Now().Month())

	var total float64

	// Meses Tener
	tener := mount - r.MountInit + 1
	if tener <= 0 {
		tener = 1
	}

	// Meses pagados
	pagadas := r.MountsToPay * r.Complete

	total = (tener - pagadas) * (r.PriceDay() * 30)

	return total
}

func (r *Resumen) PriceDay() float64 {
	return r.PriceYear / (DAYS_MOUNTH * MOUNTHS_YEAR)
}

func (r *Resumen) PriceDayNow() float64 {
	today := time.Now()
	daysOfMountsToday := float64(today.Month()) * DAYS_MOUNTH
	dayToday := today.YearDay()

	return r.PayMountNow() / (daysOfMountsToday - float64(dayToday))
}

func (r *Resumen) Resumen(salary float64) []string {
	info := make([]string, 5)

	info[0] = r.Type
	info[1] = r.Name
	info[2] = fmt.Sprintf("%%%.2f", (r.PriceDayNow()/240)*100)
	info[3] = fmt.Sprintf("$%.2f", r.PriceDayNow())
	info[4] = fmt.Sprintf("%.2f", r.PayMountNow())

	return info
}
