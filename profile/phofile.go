package profile

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/vicmanbrile/moneyGolang/profile/expenses"
	"github.com/vicmanbrile/moneyGolang/profile/status"
)

type Perfil struct {
	Creditos     []expenses.Credit      `json:"credit"`
	Deudas       []expenses.Debt        `json:"debts"`
	Suscriptions []expenses.Suscription `json:"suscriptions"`
	Percentiles  []expenses.Percentile  `json:"percentile"`
	Registers    status.Registers       `json:"registers"`
	Wallets      status.Wallet          `json:"wallets"`
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

func (p *Perfil) PrintTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetBorder(false)

	table.SetHeader([]string{"Grupo", "Descripcion", "Porcentaje", "$ X D"})

	info := make([][]string, 0)
	{
		for _, value := range p.CalcPerfil() {
			d := value.Resumen(p.Wallets.Average)
			info = append(info, d)
		}
	}

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

	total := p.Registers.Budgets()

	fmt.Println(total)
	fmt.Println(total.Free((p.PriceDays() / p.Wallets.Average), &p.Wallets))

	/*
		table := tablewriter.NewWriter(os.Stdout)
		table.SetBorder(false)

		table.SetHeader([]string{"Tiempo", "Dias", "Debemos", "Libres", "Falta", "Guardados", "Extras"})

		info := make([][]string, 0)
		{
			var BudgetWon []string

			b := p.Registers.Budgets()

			BudgetWon = append(BudgetWon, "Ganado")
			BudgetWon = append(BudgetWon, fmt.Sprintf("%.2f", b.Must(p.PriceDays()/p.Wallets.Average)))
			BudgetWon = append(BudgetWon, fmt.Sprintf("%.2f", b.Free(p.PriceDays()/p.Wallets.Average)))
			BudgetWon = append(BudgetWon, fmt.Sprintf("%.2f", b.Lack()))

			info = append(info, BudgetWon)
		}

		table.SetFooter([]string{
			"",
			"",
			"",
			"",
		})

		for _, v := range info {
			table.Append(v)
		}

		table.Render()
	*/
}
