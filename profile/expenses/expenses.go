package expenses

import (
	"time"

	"github.com/vicmanbrile/moneyGolang/dates"
)

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
				r.MonthFinish = float64(c.Datails.Mensualy)
			}
		case "Debt":
			{
				r.MonthFinish = dates.Today.Mounth()
			}
		case "Percentile":
			{
				r.MonthFinish = dates.MOUNTHS_YEAR

				{
					var Procintile = c.Datails.Optionals.Percentage * dates.DAYS_YEAR

					c.Datails.Precing = Procintile * Average
				}
			}
		case "Suscription":
			{
				r.MonthFinish = dates.MOUNTHS_YEAR
				switch c.Datails.Optionals.Suscription {
				case "yearly":
					{
						return
					}
				case "monthly":
					{
						c.Datails.Precing *= dates.MOUNTHS_YEAR
					}
				}
			}
		}
	}

	/* Definimos el precio por dias de los creditos */
	if c.Datails.Interes > 0 {
		c.Datails.Precing *= (c.Datails.Interes + 1)
	}

	r.Price = dates.ToPriceInDays(c.Datails.Precing, Average)

	if c.Spent == 0 {
		r.Paid = 0
	} else {
		r.Paid = dates.ToPriceInDays(c.Spent, Average)
	}

	r.Subtrac = r.Price - r.Paid
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

type Resumen struct {
	Name        string
	Type        string
	Price       dates.PriceInDays
	Paid        dates.PriceInDays
	Subtrac     dates.PriceInDays
	MonthFinish float64

	// Tipos nuevos para el calendario
	DateInit   time.Time
	DateFinish time.Time
}

func (r *Resumen) PorcentileForYear() float64 {
	return float64(r.Price) / dates.DAYS_YEAR
}
