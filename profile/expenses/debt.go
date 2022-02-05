package expenses

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

	r.PriceYear = d.Amount

	r.Porcentile = (r.PriceYear / salary) / (DAYS_MOUNTH * MOUNTHS_YEAR)

	r.Complete = d.Spent / r.PriceYear

	return r
}
