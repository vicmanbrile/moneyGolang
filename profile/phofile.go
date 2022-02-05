package profile

import (
	"fmt"

	"github.com/vicmanbrile/moneyGolang/profile/expenses"
	"github.com/vicmanbrile/moneyGolang/profile/status"
)

type Perfil struct {
	Expenses  expenses.Expenses `json:"expenses"`
	Registers status.Registers  `json:"registers"`
	Wallets   status.Wallet     `json:"wallets"`
}

func (p *Perfil) StutusTable() {

	total := p.Registers.Budgets()

	fmt.Printf("%+v\n", total)
	fmt.Println(total.Free((p.Expenses.PriceDays(p.Wallets.Average) / p.Wallets.Average), &p.Wallets))
}
