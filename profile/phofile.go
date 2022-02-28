package profile

import (
	"fmt"
	"time"
)

type Perfil struct {
	Registers Registers `json:"registers"`
	Wallets   Wallet    `json:"wallets"`
}

func (p *Perfil) Free() string {

	porcentile := p.Wallets.Expenses.CalcPerfil(p.Wallets.Average)

	diasPagadas := p.Registers.Budgets().Entries / p.Wallets.Average

	debemos := p.Wallets.Average * porcentile.FaltaMount() * diasPagadas

	tenemos := p.Wallets.Total() + (porcentile.PriceDays() * float64(time.Now().Month()) * 30 * p.Wallets.Average)

	libres := tenemos - debemos

	result := fmt.Sprintf(`

	Dias Pagadas: %.2f
	Now: %.0f

	Debemos: %.2f

	Tenemos: %.2f

	Libres: %.2f
	`, diasPagadas, automaticTime(), debemos, tenemos, libres)

	return result
}

func automaticTime() float64 {
	today := time.Now().YearDay()

	return float64(today)
}
