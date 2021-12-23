package main

type Suscription struct {
	Name    string  `json:"name"`
	Type    string  `json:"type"`
	Pricing float64 `json:"pricing"`
}

func (s Suscription) PriceMount() float64 {
	switch s.Type {
	case "monthly":
		return s.Pricing
	case "yearly":
		return s.Pricing / float64(12)
	default:
		return 0
	}
}
