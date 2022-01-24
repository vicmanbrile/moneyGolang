package status

import (
	"fmt"
	"time"
)

type Registers struct {
	Spent []struct {
		Key   string  `json:"key"`
		Value float64 `json:"value"`
	} `json:"spent"`
	Entries []struct {
		Week  float64 `json:"week"`
		Money float64 `json:"money"`
	} `json:"entries"`
}

func (r *Registers) Budgets() (Bdgt Budget) {
	for _, value := range r.Entries {
		Bdgt.Entries += value.Money
	}

	for _, value := range r.Spent {
		Bdgt.Spent += value.Value
	}

	return
}

type Budget struct {
	Spent   float64
	Entries float64
}

func (b *Budget) Free(percentage float64, w *Wallet) string {

	debemos := w.Average * percentage * automaticTime()
	extras := (((b.Entries / w.Average) - automaticTime()) * w.Average)

	tenemos := w.Total()
	pagado := b.Spent

	var libres float64

	if extras >= 0 {
		libres = (tenemos + pagado) - (debemos + extras)
	} else {
		libres = (tenemos + pagado) - (debemos - extras)
	}

	result := fmt.Sprintf(`
	Debemos: %.2f
	Extras: %.2f

	Tenemos: %.2f
	Pagado: %.2f

	Libres: %.2f
	`, debemos, extras, tenemos, pagado, libres)

	return result
}

func automaticTime() float64 {
	today := time.Now().YearDay()

	return float64(today)
}
