package profile

import "github.com/vicmanbrile/moneyGolang/profile/expenses"

type Wallet struct {
	Cash     float64           `json:"cash"`
	Banking  float64           `json:"banking"`
	Average  float64           `json:"average"`
	Expenses expenses.Expenses `json:"expenses"`
}

func (w *Wallet) Total() float64 {
	total := w.Cash + w.Banking

	return total
}

func (w *Wallet) BudgetsDays() float64 {
	return w.Total() / w.Average
}
