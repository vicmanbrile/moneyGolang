package db

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID primitive.ObjectID
}

func (u *User) ReadDeposit(mc *MongoConnection) (rest *Deposits) {
	bsonData, err := mc.FindOne(u.ID, "deposits")

	if err != nil {
		fmt.Println(err)
	}

	err = bson.Unmarshal(bsonData, rest)
	if err != nil {
		fmt.Println(err)
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

type Wallet struct {
	Cash    float64
	Bank    float64
	Savings float64
}

type Suscriptions struct{}
