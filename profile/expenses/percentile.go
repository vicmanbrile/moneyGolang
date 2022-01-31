package expenses

import "time"

type Percentile struct {
	Name       string  `json:"name"`
	Percentage float64 `json:"Percentage"`
	Spent      float64 `json:"spent"`
}

func (p *Percentile) CalcPercentiles(salary float64) *Resumen {
	var r = &Resumen{
		Name: p.Name,
		Type: "Porcentil",
	}

	dayNow := time.Now().AddDate(0, 0, 30).YearDay()

	r.PriceDay = ((p.Percentage * salary * float64(dayNow)) - p.Spent) / float64(DAYS_MOUNTH)
	r.ProcentileFloat = r.PriceDay / salary

	return r
}
