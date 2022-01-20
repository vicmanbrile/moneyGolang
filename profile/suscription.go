package profile

type Suscription struct {
	Name    string  `json:"name"`
	Type    string  `json:"type"`
	Pricing float64 `json:"pricing"`
}

func (s *Suscription) CalcSuscriptions() *Resumen {
	var r = &Resumen{
		Name: s.Name,
		Type: "Mensualidad",
	}
	switch s.Type {
	case "yearly":
		r.PriceMount = s.Pricing / MOUNTHS_YEAR
	default:
		r.PriceMount = s.Pricing
	}

	return r
}
