package main

import "fmt"

type Suscription struct {
	Name    string  `json:"name"`
	Type    string  `json:"type"`
	Pricing float64 `json:"pricing"`
}

func (s Suscription) PriceMount() float64 {
	switch s.Type {
	case "yearly":
		return s.Pricing / MOUNTHS_YEAR
	default:
		return s.Pricing
	}
}

func (s Suscription) PriceForDays() float64 {
	return s.PriceMount() / DAYS_MOUNTH
}

func (s Suscription) Resumen(salary float64) []string {
	info := make([]string, 4)

	info[0] = "Suscripciones"
	info[1] = s.Name
	info[2] = fmt.Sprintf("%%%.2f", (s.PriceForDays()/salary)*100)
	info[3] = fmt.Sprintf("$%.2f", s.PriceForDays())

	return info
}
