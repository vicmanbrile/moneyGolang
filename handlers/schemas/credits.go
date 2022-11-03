package schemas

import (
	"time"

	"github.com/vicmanbrile/moneyGolang/application/dates"
)

type Resumen struct {
	Name    string            `json:"name"`
	Price   dates.PriceInDays `json:"price"`
	Paid    dates.PriceInDays `json:"paid"`
	Subtrac dates.PriceInDays `json:"subtrac"`

	// Tipos nuevos para el calendario
	DateInit   time.Time `json:"date_init"`
	DateFinish time.Time `json:"date_finish"`
}

type MostrarDeposits struct {
	Average  float64 `json:"average"`
	YearDay  int     `json:"year_day"`
	Deposits float64 `json:"deposits"`
}
