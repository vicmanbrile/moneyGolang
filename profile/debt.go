package profile

import "fmt"

type Debt struct {
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
	Days   float64 `json:"days"`
}

func (d Debt) PriceMount() float64 {
	total := (d.Amount / d.Days) * DAYS_MOUNTH
	return float64(total)
}

func (d Debt) PriceForDays() float64 {
	return d.PriceMount() / DAYS_MOUNTH
}

func (d Debt) Resumen(salary float64) []string {
	info := make([]string, 4)

	info[0] = "Deudas"
	info[1] = d.Name
	info[2] = fmt.Sprintf("%.2f%%", (d.PriceForDays()/salary)*100)
	info[3] = fmt.Sprintf("$%.2f", d.PriceForDays())

	return info
}
