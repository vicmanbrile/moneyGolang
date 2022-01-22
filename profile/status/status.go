package status

import (
	"time"
)

type Registers struct {
	Spent []struct {
		Key   string  `json:"key"`
		Value float64 `json:"value"`
	} `json:"spent"`
	Saved []struct {
		Key   string  `json:"key"`
		Value float64 `json:"value"`
	} `json:"saved"`
	Extras []struct {
		Week  float64 `json:"week"`
		Extra float64 `json:"extra"`
		Days  float64 `json:"days"`
	} `json:"extras"`
}

type Budget struct {
	Total    float64
	Restados float64
	Dias     float64
}

func (r *Registers) BudgetsNow(money float64) (Bdgt Budget) {
	Bdgt.Dias = automaticTime()
	Bdgt.Total = Bdgt.Dias * money

	{
		var restar float64
		for _, value := range r.Saved {
			restar += value.Value
		}

		for _, value := range r.Spent {
			restar += value.Value
		}

		Bdgt.Restados = restar
	}

	return
}

func (r *Registers) BudgetsWon(money float64) (Bdgt Budget) {
	var days float64
	for _, value := range r.Extras {
		days += value.Days
	}
	{
		var restar float64
		for _, value := range r.Saved {
			restar += value.Value
		}

		for _, value := range r.Spent {
			restar += value.Value
		}

		Bdgt.Restados = restar
	}

	Bdgt.Dias = days
	Bdgt.Total = days * money
	return
}

func (b *Budget) Lack(percentage float64) float64 {
	return (b.Total * percentage) - b.Restados
}

func (b *Budget) Free(percentage float64) float64 {
	return b.Total * (1 - percentage)
}

func automaticTime() float64 {
	today := time.Now().YearDay()

	return float64(today)
}
