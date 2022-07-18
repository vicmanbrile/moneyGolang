package profile

import "github.com/vicmanbrile/moneyGolang/app/expenses"

type Wallet struct {
	Cash     float64           `bson:"cash"`
	Banking  float64           `bson:"banking"`
	Average  float64           `bson:"average"`
	Expenses expenses.Expenses `bson:"expenses"`
}

func (w *Wallet) Total() float64 {
	total := w.Cash + w.Banking

	return total
}

func (w *Wallet) BudgetsDays() float64 {
	return w.Total() / w.Average
}
