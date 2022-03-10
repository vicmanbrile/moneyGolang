package profile

type Perfil struct {
	Registers Registers `json:"registers"`
	Wallets   Wallet    `json:"wallets"`
}
