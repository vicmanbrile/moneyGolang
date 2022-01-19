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
	Total float64
}

func (r *Registers) BudgetsNow(m float64) (Bdgt Budget) {
	days := automaticTime()
	Bdgt.Total = days * m
	return
}

func (r *Registers) BudgetsWon(m float64) (Bdgt Budget) {
	var days float64
	for _, value := range r.Extras {
		days += value.Days
	}

	Bdgt.Total = days * m
	return
}

func (b *Budget) Lack(percentage float64) float64 {
	return b.Total * percentage
}

func (b *Budget) Free(percentage float64) float64 {
	return b.Total * percentage
}

func automaticTime() float64 {
	today := time.Now().YearDay()

	return float64(today)
}
