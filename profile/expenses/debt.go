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

	r.PriceDay = (d.Amount - d.Spent) / float64(DAYS_MOUNTH)
	r.ProcentileFloat = r.PriceDay / salary

	return r
}
