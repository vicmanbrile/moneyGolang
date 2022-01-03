package main

var (
	DAYS_MOUNTH  float64 = 30
	MOUNTHS_YEAR float64 = 12
)

type Perfil struct {
	Creditos     []Product     `json:"credit"`
	Deudas       []Debt        `json:"debts"`
	Suscriptions []Suscription `json:"suscriptions"`
	Wallets      Wallet        `json:"wallets"`
	Percentiles  []Percentile  `json:"percentile"`
}
