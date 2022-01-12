package profile

import "fmt"

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

func (p Credit) PriceMount() float64 {
	var price float64
	if p.Datails.Interes > 0 && p.Datails.Interes < 1 {
		price = p.Datails.Precing * (p.Datails.Interes + 1)
	} else {
		price = p.Datails.Precing
	}

	if p.Datails.Mensualy == 0 {
		return 0
	} else {
		return price / float64(p.Datails.Mensualy)
	}

}

func (p Credit) PriceForDays() float64 {
	return p.PriceMount() / DAYS_MOUNTH
}

func (p Credit) Resumen(salary float64) []string {
	info := make([]string, 4)

	info[0] = "Credito"
	info[1] = p.Name
	info[2] = fmt.Sprintf("%.2f%%", (p.PriceForDays()/salary)*100)
	info[3] = fmt.Sprintf("$%.2f", p.PriceForDays())

	return info
}

func (p Credit) UltimeMount() int {
	return (p.Date.Mount + p.Datails.Mensualy) % 12
}

func (p Credit) UltimeYear() int {
	return p.Date.Year + (p.Date.Mount/12 + p.Datails.Mensualy/12)
}
