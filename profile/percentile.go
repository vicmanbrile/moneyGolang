package profile

type Percentile struct {
	Name       string  `json:"name"`
	Percentage float64 `json:"Percentage"`
}

func (p *Percentile) CalcPercentiles(salary float64) *Resumen {
	var r = &Resumen{
		Name: p.Name,
		Type: "Porcentil",
	}

	r.PriceMount = p.Percentage * DAYS_MOUNTH * salary
	return r
}
