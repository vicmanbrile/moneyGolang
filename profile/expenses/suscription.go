package expenses

import "time"

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

	dayNow := time.Now().AddDate(0, 0, 30).YearDay() / DAYS_MOUNTH

	switch s.Type {
	case "yearly":
		r.PriceDay = (((s.Pricing / float64(MOUNTHS_YEAR)) * float64(dayNow)) - s.Spent) / float64(DAYS_MOUNTH)
	default:
		r.PriceDay = ((s.Pricing * float64(dayNow)) - s.Spent) / float64(DAYS_MOUNTH)
	}

	r.ProcentileFloat = (r.PriceDay) / salary

	return r
}
