package expenses

import (
	"time"

	"github.com/vicmanbrile/moneyGolang/app/dates"
	"github.com/vicmanbrile/moneyGolang/schemas"
)

type Credits struct {
	Name string `bson:"name"`
	Type string `bson:"type"`
	Date struct {
		Mount int `bson:"mount"`
		Year  int `bson:"year"`
	} `bson:"date"`
	Datails struct {
		Interes  float64 `bson:"interes"`
		Precing  float64 `bson:"precing"`
		Mensualy int     `bson:"mensualy"`
	} `bson:"datails"`
	Spent float64 `bson:"spent"`
}

func (c *Credits) Calculator(Average float64) (r schemas.Resumen) {
	r = schemas.Resumen{
		Name: c.Name,
	}

	{ // Establecer fecha de inicio y final

		{ // Inicio
			r.DateInit = time.Date(c.Date.Year, time.Month(c.Date.Mount), 1, 0, 0, 0, 0, time.UTC)
		}

		{ // Final

			DaysMount := dates.DayOfYear((c.Datails.Mensualy * int(dates.DAYS_MOUNTH))).ToFraction()

			r.DateFinish = r.DateInit.AddDate(int(DaysMount.Years), int(DaysMount.Mounts), int(DaysMount.Days))
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
	Creditos []Credits `bson:"credit"`
}

func (e *Expenses) CalcPerfil(Average float64) (TR []schemas.Resumen) {
	for _, value := range e.Creditos {
		TR = append(TR, value.Calculator(Average))
	}

	return
}
