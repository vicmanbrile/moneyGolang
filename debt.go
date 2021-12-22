package main

type Debt struct {
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
	Days   float64 `json:"days"`
}

func (d Debt) PriceMount() float64 {
	total := (d.Amount / d.Days) * 30
	return float64(total)
}
