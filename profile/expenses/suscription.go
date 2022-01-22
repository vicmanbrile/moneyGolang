package expenses

type Suscription struct {
	Name    string  `json:"name"`
	Type    string  `json:"type"`
	Pricing float64 `json:"pricing"`
}

func (s *Suscription) CalcSuscriptions(salary float64) *Resumen {
	var r = &Resumen{
		Name: s.Name,
		Type: "Mensualidad",
	}
	switch s.Type {
	case "yearly":
		r.PriceDay = s.Pricing / MOUNTHS_YEAR / DAYS_MOUNTH
	default:
		r.PriceDay = s.Pricing / DAYS_MOUNTH
	}

	r.ProcentileFloat = r.PriceDay / salary

	return r
}
