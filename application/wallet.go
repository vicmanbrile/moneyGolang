package application

type Wallet struct {
	Cash    float64
	Bank    float64
	Savings float64
}

func (w *Wallet) Total() float64 {
	total := w.Cash + w.Bank

	return total
}

func (w *Wallet) BudgetsDays(d Deposits) float64 {
	return w.Total() / d.Average()
}
