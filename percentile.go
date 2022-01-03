package main

import "fmt"

type Percentile struct {
	Name       string  `json:"name"`
	Percentage float64 `json:"Percentage"`
}

func (p Percentile) PriceMount(salary float64) float64 {
	total := salary * p.Percentage * DAYS_MOUNTH
	return total
}

func (p Percentile) PriceForDays(salary float64) float64 {
	total := salary * p.Percentage
	return total
}

func (p Percentile) Resumen(salary float64) []string {
	info := make([]string, 4)

	info[0] = "Porcentil"
	info[1] = p.Name
	info[2] = fmt.Sprintf("%.2f%%", (p.PriceForDays(salary)/salary)*100)
	info[3] = fmt.Sprintf("$%.2f", p.PriceForDays(salary))

	return info
}
