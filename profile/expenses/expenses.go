package expenses

import (
	"fmt"
	"math"
	"os"
	"time"

	"github.com/olekukonko/tablewriter"
)

var (
	DAYS_WEEK    float64 = 10
	DAYS_MOUNTH  float64 = DAYS_WEEK * 3
	MOUNTHS_YEAR float64 = 12
	DAYS_YEAR    float64 = DAYS_MOUNTH * MOUNTHS_YEAR
)

type DayOfYear float64

func (dfy DayOfYear) Mounth() float64 {
	return math.Ceil(float64(dfy) / DAYS_MOUNTH)
}

func (dfy DayOfYear) Week() float64 {
	return math.Ceil(float64(dfy) / DAYS_WEEK)
}

var (
	Today DayOfYear = DayOfYear(time.Now().YearDay())
)

type Expenses struct {
	Creditos []Credits `json:"credit"`
}
type Credits struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Date struct {
		Mount int `json:"mount"`
		Year  int `json:"year"`
	} `json:"date"`
	Datails struct {
		Interes   float64 `json:"interes"`
		Precing   float64 `json:"precing"`
		Mensualy  int     `json:"mensualy"`
		Optionals struct {
			Percentage  float64 `json:"porcentile"`
			Suscription string  `json:"suscription"`
		} `json:"optionals"`
	} `json:"datails"`
	Spent float64 `json:"spent"`
}

func (c *Credits) Calculator(Average float64) (r Resumen) {
	r = Resumen{
		Name: c.Name,
		Type: c.Type,
	}

	{ /* Establecer los tiempos de pago */
		switch c.Type {
		case "Credit":
			{
				r.MountInit = float64(c.Date.Mount)
				r.MountsToPay = float64(c.Datails.Mensualy)
			}
		case "Debt":
			{
				r.MountInit = Today.Mounth()
				r.MountsToPay = 1
			}
		case "Percentile":
			{
				r.MountInit = 1
				r.MountsToPay = 12

				{
					var ProcintileAll = c.Datails.Optionals.Percentage + 1

					r.Price = PriceInDays(ProcintileAll * Average)
				}
			}
		case "Suscription":
			{
				switch c.Datails.Optionals.Suscription {
				case "yearly":
					r.Price = PriceInDays(c.Datails.Precing)

				case "monthly":
					{
						r.Price = PriceInDays(c.Datails.Precing * MOUNTHS_YEAR)
					}
				}
			}
		}
	}

	/* Definimos el precio por dias de los creditos */
	if c.Datails.Interes > 0 {
		c.Datails.Precing *= c.Datails.Interes + 1
	}

	r.Price = ToPriceInDays(c.Datails.Precing, Average)
	r.Paid = PriceInDays(c.Spent / Average)

	return
}

func ToPriceInDays(Money float64, Average float64) (PD PriceInDays) {
	return PriceInDays(Money / Average)
}

func (e *Expenses) CalcPerfil(Average float64) AllExpenses {
	var AE = AllExpenses{}

	for _, value := range e.Creditos {
		AE.ToDoExpenses = append(AE.ToDoExpenses, value.Calculator(Average))
	}

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

type Calculator interface {
	CalculatorResumen(Average float64) Resumen
}

func CalculatorResumen(w Calculator, Average float64) Resumen {
	return w.CalculatorResumen(Average)
}

type PriceInDays float64

func (PID *PriceInDays) Week() {

}

type Resumen struct {
	Name        string
	Type        string
	Price       PriceInDays
	Paid        PriceInDays
	MountInit   float64
	MountsToPay float64
}

func (r *Resumen) PorcentileNow() (porcentaje float64) {

	// Dias del año
	PriceInDays := (float64(r.Price) / 360) * DAYS_YEAR

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

	priceDaysComplete := r.Paid * (r.Price / r.Paid) * PriceInDays(DAYS_YEAR)

	return (float64(priceDaysComplete) / priceToDays) * r.PorcentileNow()
}

func (r *Resumen) Resumen() []string {
	info := make([]string, 4)

	info[0] = r.Type
	info[1] = r.Name
	info[2] = fmt.Sprintf("%.2f%%", r.PorcentileNow()*100)
	info[3] = fmt.Sprintf("%.2f%%", r.PorcentileComplete()*100)

	return info
}
