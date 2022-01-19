package profile

type Credit struct {
	Name string `json:"name"`
	Date struct {
		Mount int `json:"mount"`
		Year  int `json:"year"`
	} `json:"date"`
	Datails struct {
		Interes  float64 `json:"interes"`
		Precing  float64 `json:"precing"`
		Mensualy int     `json:"mensualy"`
	} `json:"datails"`
}

func (c *Credit) CalcCredit() *Resumen {
	r := &Resumen{
		Name: c.Name,
		Type: "Credito",
	}
	{
		var price float64
		if c.Datails.Interes > 0 && c.Datails.Interes < 1 {
			price = c.Datails.Precing * (c.Datails.Interes + 1)
		} else {
			price = c.Datails.Precing
		}

		if c.Datails.Mensualy == 0 {
			r.PriceMount = price
		} else {
			r.PriceMount = price / float64(c.Datails.Mensualy)
		}
	}

	return r
}
