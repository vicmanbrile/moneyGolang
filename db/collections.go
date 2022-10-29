package db

import "go.mongodb.org/mongo-driver/bson/primitive"

type Profile struct {
	ID        primitive.ObjectID
	Deposits  primitive.ObjectID
	Shoppings primitive.ObjectID
	Wallet    primitive.ObjectID
}

type Deposits struct {
	YearDay  int
	Deposits float64
}

func (d *Deposits) Read() {}

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

type Wallet struct {
	Cash    float64
	Bank    float64
	Savings float64
}
