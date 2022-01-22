package expenses

type Debt struct {
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
	Days   float64 `json:"days"`
}

func (d *Debt) CalcDebt(salary float64) *Resumen {
	r := &Resumen{
		Name: d.Name,
		Type: "Deuda",
	}

	r.PriceDay = d.Amount / d.Days
	r.ProcentileFloat = r.PriceDay / salary

	return r
}
