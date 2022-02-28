package expenses

import (
	"fmt"
	"math"
	"os"
	"time"

	"github.com/olekukonko/tablewriter"
)

var (
	DAY          float64 = 1
	DAYS_WEEK    float64 = DAY * 10
	DAYS_MOUNTH  float64 = DAYS_WEEK * 3
	MOUNT        float64 = 1
	MOUNTHS_YEAR float64 = 12
	DAYS_YEAR    float64 = DAYS_MOUNTH * MOUNTHS_YEAR
)

func ToPriceInDays(Money float64, Average float64) PriceInDays {
	return PriceInDays(Money / Average)
}

var (
	Today DayOfYear = DayOfYear(time.Now().YearDay())
)

type DayOfYear float64

func (dfy DayOfYear) Mounth() float64 {
	return math.Ceil(float64(dfy) / DAYS_MOUNTH)
}

func (dfy DayOfYear) Week() float64 {
	return math.Ceil(float64(dfy) / DAYS_WEEK)
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
				r.MountsToPay = MOUNT
			}
		case "Percentile":
			{
				r.MountInit = MOUNT
				r.MountsToPay = MOUNTHS_YEAR

				{
					var Procintile = c.Datails.Optionals.Percentage * DAYS_YEAR

					c.Datails.Precing = Procintile * Average
				}
			}
		case "Suscription":
			{
				r.MountInit = float64(c.Date.Mount)
				switch c.Datails.Optionals.Suscription {
				case "yearly":
					{
						r.MountsToPay = MOUNTHS_YEAR
					}
				case "monthly":
					{
						r.MountsToPay = MOUNT
					}
				}
			}
		}
	}

	/* Definimos el precio por dias de los creditos */
	if c.Datails.Interes > 0 {
		c.Datails.Precing *= (c.Datails.Interes + 1)
	}

	r.Price = ToPriceInDays(c.Datails.Precing, Average)

	if c.Spent == 0 {
		r.Paid = 0
	} else {
		r.Paid = ToPriceInDays(c.Spent, Average)
	}

	return
}

type Expenses struct {
	Creditos []Credits `json:"credit"`
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

	// Porcentaje a pagar por mes
	PorcentajeForMount := 1 / r.MountsToPay

	// Meses en total para ahora
	Geting := (Today.Mounth() + 1) - r.MountInit
	if Geting <= 0 {
		Geting = 1
	}

	// --
	PorcentileSaved := PorcentajeForMount * Geting

	formula := (float64(r.Price) * PorcentileSaved) / (Today.Mounth() * DAYS_MOUNTH)

	return formula
}

func (r *Resumen) PorcentileComplete() float64 {
	if r.Paid == 0 {
		return 0
	} else {
		return (float64(r.Paid) / (Today.Mounth() * DAYS_MOUNTH)) / r.PorcentileNow()
	}
}

func (r *Resumen) Resumen() []string {
	info := make([]string, 4)

	info[0] = r.Type
	info[1] = r.Name
	info[2] = fmt.Sprintf("%.2f%%", r.PorcentileNow()*100)
	info[3] = fmt.Sprintf("%.2f%%", r.PorcentileComplete()*100)

	return info
}
