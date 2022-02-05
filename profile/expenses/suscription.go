package expenses

type Suscription struct {
	Name    string  `json:"name"`
	Type    string  `json:"type"`
	Pricing float64 `json:"pricing"`
	Spent   float64 `json:"spent"`
}

func (s *Suscription) CalcSuscriptions(salary float64) *Resumen {
	var r = &Resumen{
		Name: s.Name,
		Type: "Mensualidad",
	}

	switch s.Type {
	case "yearly":
		r.PriceYear = s.Pricing
	default:
		r.PriceYear = s.Pricing * MOUNTHS_YEAR
	}

	r.Porcentile = (r.PriceYear / salary) / (DAYS_MOUNTH * MOUNTHS_YEAR)

	r.Complete = s.Spent / r.PriceYear

	return r
}
