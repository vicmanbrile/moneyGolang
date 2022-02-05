package expenses

import "fmt"

var (
	DAYS_MOUNTH  float64 = 30
	MOUNTHS_YEAR float64 = 12
)

type Resumen struct {
	PriceYear  float64
	Porcentile float64
	Complete   float64
	Name       string
	Type       string
}

func (r *Resumen) PriceDay() float64 {
	return r.PriceYear / (DAYS_MOUNTH * MOUNTHS_YEAR)
}

func (r *Resumen) Resumen(salary float64) []string {
	info := make([]string, 5)

	info[0] = r.Type
	info[1] = r.Name
	info[2] = fmt.Sprintf("%%%.2f", r.Porcentile*100)
	info[3] = fmt.Sprintf("$%.2f", r.PriceDay())
	info[4] = fmt.Sprintf("%%%.2f", r.Complete*100)

	return info
}
