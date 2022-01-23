package profile

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/vicmanbrile/moneyGolang/profile/expenses"
	"github.com/vicmanbrile/moneyGolang/profile/status"
)

type Perfil struct {
	Wallets      Wallet                 `json:"wallets"`
	Creditos     []expenses.Credit      `json:"credit"`
	Deudas       []expenses.Debt        `json:"debts"`
	Suscriptions []expenses.Suscription `json:"suscriptions"`
	Percentiles  []expenses.Percentile  `json:"percentile"`
	Registers    status.Registers       `json:"registers"`
}

func (p *Perfil) CalcPerfil() []expenses.Resumen {
	var Todos []expenses.Resumen

	var mySalary float64 = p.Wallets.Average

	for _, value := range p.Creditos {
		Todos = append(Todos, *value.CalcCredit(mySalary))
	}
	for _, value := range p.Deudas {
		Todos = append(Todos, *value.CalcDebt(mySalary))
	}

	for _, value := range p.Suscriptions {
		Todos = append(Todos, *value.CalcSuscriptions(mySalary))
	}

	for _, value := range p.Percentiles {
		Todos = append(Todos, *value.CalcPercentiles(p.Wallets.Average))
	}

	return Todos
}

func (p *Perfil) PriceDays() float64 {
	var result float64

	for _, value := range p.CalcPerfil() {
		result += value.PriceDay
	}

	return result
}

func (p *Perfil) Resumen() [][]string {
	info := make([][]string, 0)

	for _, value := range p.CalcPerfil() {
		d := value.Resumen(p.Wallets.Average)
		info = append(info, d)
	}

	return info
}

func (p *Perfil) PrintTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetBorder(false)

	table.SetHeader([]string{"Grupo", "Descripcion", "Porcentaje", "$ X D"})

	info := p.Resumen()

	table.SetFooter([]string{
		"",
		"Total:",
		fmt.Sprintf("%.2f%%", (p.PriceDays()/p.Wallets.Average)*100),
		fmt.Sprintf("$%.2f", p.PriceDays()),
	})

	for _, v := range info {
		table.Append(v)
	}

	table.Render()
}

func (p *Perfil) StutusTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetBorder(false)

	table.SetHeader([]string{"Tiempo", "Dias", "Debemos", "Libres", "Falta", "Guardados", "Extras"})

	info := p.StatusResumen()

	table.SetFooter([]string{
		"",
		"",
		"",
		"",
		"",
		"",
		"",
	})

	for _, v := range info {
		table.Append(v)
	}

	table.Render()
}

func (p *Perfil) StatusResumen() [][]string {
	info := make([][]string, 0)

	{
		var BudgetNow []string

		b := p.Registers.BudgetsNow(p.Wallets.Average)

		BudgetNow = append(BudgetNow, "Actual")
		BudgetNow = append(BudgetNow, fmt.Sprintf("%.0f", b.Dias))
		BudgetNow = append(BudgetNow, fmt.Sprintf("%.2f", b.Must(p.PriceDays()/p.Wallets.Average)))
		BudgetNow = append(BudgetNow, fmt.Sprintf("%.2f", b.Free(p.PriceDays()/p.Wallets.Average)))
		BudgetNow = append(BudgetNow, fmt.Sprintf("%.2f", b.Lack()))
		BudgetNow = append(BudgetNow, fmt.Sprintf("%.2f", b.SaveNSpent))
		BudgetNow = append(BudgetNow, fmt.Sprintf("%.2f", b.Extra))

		info = append(info, BudgetNow)
	}

	{
		var BudgetWon []string

		b := p.Registers.BudgetsWon(p.Wallets.Average)

		BudgetWon = append(BudgetWon, "Ganado")
		BudgetWon = append(BudgetWon, fmt.Sprintf("%.0f", b.Dias))
		BudgetWon = append(BudgetWon, fmt.Sprintf("%.2f", b.Must(p.PriceDays()/p.Wallets.Average)))
		BudgetWon = append(BudgetWon, fmt.Sprintf("%.2f", b.Free(p.PriceDays()/p.Wallets.Average)))
		BudgetWon = append(BudgetWon, fmt.Sprintf("%.2f", b.Lack()))
		BudgetWon = append(BudgetWon, fmt.Sprintf("%.2f", b.SaveNSpent))
		BudgetWon = append(BudgetWon, fmt.Sprintf("%.2f", b.Extra))

		info = append(info, BudgetWon)
	}

	return info
}
