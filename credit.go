package main

type Perfil struct {
	Creditos []Product `json:"credit"`
}

type Product struct {
	ID   string `json:"id"`
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

func (p Product) priceMount() float64 {
	var price float64
	if p.Datails.Interes > 0 && p.Datails.Interes < 1 {
		price = p.Datails.Precing * (p.Datails.Interes + 1)
	} else {
		price = p.Datails.Precing
	}
	return price / float64(p.Datails.Mensualy)
}

func (p Product) ultimeMount() int {
	return (p.Date.Mount + p.Datails.Mensualy) % 12
}

func (p Product) ultimeYear() int {
	return p.Date.Year + (p.Date.Mount/12 + p.Datails.Mensualy/12)
}
