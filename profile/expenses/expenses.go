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
	DAYS_YEAR    float64 = DAYS_MOUNTH * MOUNTHS_YEAR
)

type Calculator interface {
	CalculatorResumen(Average float64) Resumen
}

func CalculatorResumen(w Calculator, Average float64) Resumen {
	return w.CalculatorResumen(Average)
}

type Expenses struct {
	Creditos     []Credit      `json:"credit"`
	Deudas       []Debt        `json:"debts"`
	Suscriptions []Suscription `json:"suscriptions"`
	Percentiles  []Percentile  `json:"percentile"`
}

func (e *Expenses) CalcPerfil(Average float64) AllExpenses {
	var Todos []Resumen

	for _, value := range e.Creditos {
		Todos = append(Todos, *value.CalculatorResumen(Average))
	}
	for _, value := range e.Deudas {
		Todos = append(Todos, *value.CalculatorResumen(Average))
	}

	for _, value := range e.Suscriptions {
		Todos = append(Todos, *value.CalculatorResumen(Average))
	}

	for _, value := range e.Percentiles {
		Todos = append(Todos, *value.CalculatorResumen(Average))
	}

	var AE AllExpenses

	AE.ToDoExpenses = Todos
	return AE
}

type AllExpenses struct {
	ToDoExpenses []Resumen
}

func (e *AllExpenses) PriceDays() float64 {
	var result float64

	for _, value := range e.ToDoExpenses {
		result += value.PorcentileComplete()
	}

	return result
}

func (e *AllExpenses) FaltaMount() float64 {
	var result float64

	for _, value := range e.ToDoExpenses {
		result += value.PorcentileNow()
	}

	return result
}

func (e *AllExpenses) PrintTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetBorder(false)

	table.SetHeader([]string{"Grupo", "Descripcion", "Porcentaje", "Complete"})

	info := make([][]string, 0)
	{
		for _, value := range e.ToDoExpenses {
			d := value.Resumen()
			info = append(info, d)
		}
	}

	table.SetFooter([]string{
		"",
		"Total:",
		fmt.Sprintf("%.2f%%", e.FaltaMount()*100),
		fmt.Sprintf("%.2f%%", e.PriceDays()*100),
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

func (r *Resumen) PorcentileNow() (porcentaje float64) {

	// Dias del año
	PriceInDays := r.Porcentile * DAYS_YEAR

	// Porcenje a pagar por mes del total
	PorcentileForMount := ((13 - r.MountsToPay) / MOUNTHS_YEAR)

	MountNow := float64(time.Now().Month())
	Geting := MountNow - r.MountInit + 1
	if Geting <= 0 {
		Geting = 1
	}
	// --
	PorcentileSaved := PorcentileForMount * Geting

	formula := (PriceInDays * PorcentileSaved) / (MountNow * DAYS_MOUNTH)

	return formula
}

func (r *Resumen) PorcentileComplete() float64 {
	priceToDays := r.PorcentileNow() * (float64(time.Now().Month()) * DAYS_MOUNTH)

	priceDaysComplete := r.Complete * r.Porcentile * DAYS_YEAR

	return (priceDaysComplete / priceToDays) * r.PorcentileNow()
}

func (r *Resumen) Resumen() []string {
	info := make([]string, 4)

	info[0] = r.Type
	info[1] = r.Name
	info[2] = fmt.Sprintf("%.2f%%", r.PorcentileNow()*100)
	info[3] = fmt.Sprintf("%.2f%%", r.PorcentileComplete()*100)

	return info
}
