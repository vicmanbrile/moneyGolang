package schemas

import (
	"time"

	"github.com/vicmanbrile/moneyGolang/dates"
)

type Resumen struct {
	Name        string            `json:"name"`
	Type        string            `json:"type"`
	Price       dates.PriceInDays `json:"price"`
	Paid        dates.PriceInDays `json:"paid"`
	Subtrac     dates.PriceInDays `json:"subtrac"`
	MonthFinish float64           `json:"mount_finish"`

	// Tipos nuevos para el calendario
	DateInit   time.Time `json:"date_init"`
	DateFinish time.Time `json:"date_finish"`
}

func (r *Resumen) PorcentileForYear() float64 {
	return float64(r.Price) / dates.DAYS_YEAR
}
