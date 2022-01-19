package profile

type Debt struct {
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
	Days   float64 `json:"days"`
}

func (d *Debt) CalcDebt() *Resumen {
	r := &Resumen{
		Name: d.Name,
		Type: "Deuda",
	}

	r.PriceMount = (d.Amount / d.Days) * DAYS_MOUNTH

	return r
}
