package expenses

import "fmt"

var (
	DAYS_MOUNTH  int = 30
	MOUNTHS_YEAR int = 12
)

type Resumen struct {
	PriceDay        float64
	ProcentileFloat float64
	Name            string
	Type            string
}

func (r *Resumen) PriceForMount() float64 {
	return r.PriceDay * float64(DAYS_MOUNTH)
}

func (r *Resumen) Resumen(salary float64) []string {
	info := make([]string, 4)

	info[0] = r.Type
	info[1] = r.Name
	info[2] = fmt.Sprintf("%%%.2f", r.ProcentileFloat*100)
	info[3] = fmt.Sprintf("$%.2f", r.PriceDay)

	return info
}
