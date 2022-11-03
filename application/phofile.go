package application

import "github.com/vicmanbrile/moneyGolang/application/expenses"

type Perfil struct {
	Registers struct {
		Entries []struct {
			Week  float64 `bson:"week"`
			Money float64 `bson:"money"`
		} `bson:"entries"`
	} `bson:"registers"`
	Wallets struct {
		Cash     float64           `bson:"cash"`
		Banking  float64           `bson:"banking"`
		Average  float64           `bson:"average"`
		Expenses expenses.Expenses `bson:"expenses"`
	} `bson:"wallets"`
}

func (p *Perfil) Budgets() (Bdgt float64) {
	for _, value := range p.Registers.Entries {
		Bdgt += value.Money
	}

	return
}

type Deposits struct {
	YearDay  int     `bson:"year_day"`
	Deposits float64 `bson:"deposits"`
}

func (d *Deposits) Update(up struct {
	YearDayNow int
	NewDeposit float64
}) {
	d.YearDay += up.YearDayNow - d.YearDay
	d.Deposits += up.NewDeposit
}

func (d *Deposits) Average() float64 {
	return d.Deposits / float64(d.YearDay)
}

type Shoppings struct {
	Description string
	Precing     float64
	Interes     float64
	Date        struct {
		Mount int
		Year  int
	}
	Mensualy int
	Spent    float64
}

type Suscriptions struct{}
