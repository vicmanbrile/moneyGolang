package expenses

import "time"

type Debt struct {
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
	Spent  float64 `json:"spent"`
}

func (d *Debt) CalcDebt(salary float64) *Resumen {
	r := &Resumen{
		Name: d.Name,
		Type: "Deuda",
	}

	r.MountInit = float64(time.Now().Month())
	r.MountsToPay = 1

	r.Complete = d.Spent / d.Amount

	if r.Complete >= 1 {
		r.PriceYear = d.Amount
	} else {
		r.PriceYear = d.Amount * 12
	}

	r.Porcentile = (r.PriceYear / salary) / (DAYS_MOUNTH * MOUNTHS_YEAR)

	return r
}
