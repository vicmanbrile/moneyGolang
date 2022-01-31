package status

import (
	"fmt"
	"time"
)

type Registers struct {
	Entries []struct {
		Week  float64 `json:"week"`
		Money float64 `json:"money"`
	} `json:"entries"`
}

func (r *Registers) Budgets() (Bdgt Budget) {
	for _, value := range r.Entries {
		Bdgt.Entries += value.Money
	}

	return
}

type Budget struct {
	Spent   float64
	Entries float64
}

func (b *Budget) Free(percentage float64, w *Wallet) string {

	diasPagadas := b.Entries / w.Average

	debemos := w.Average * percentage * diasPagadas

	tenemos := w.Total()

	libres := tenemos - debemos

	result := fmt.Sprintf(`

	Dias Pagadas: %.2f
	Now: %.0f

	Debemos: %.2f

	Tenemos: %.2f

	Libres: %.2f
	`, diasPagadas, automaticTime(), debemos, tenemos, libres)

	return result
}

func automaticTime() float64 {
	today := time.Now().YearDay()

	return float64(today)
}
