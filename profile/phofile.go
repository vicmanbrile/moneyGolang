package profile

type Perfil struct {
	Registers Registers `bson:"registers"`
	Wallets   Wallet    `bson:"wallets"`
}
