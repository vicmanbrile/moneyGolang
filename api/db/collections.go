package db

import (
	"fmt"

	"github.com/vicmanbrile/moneyGolang/application"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID primitive.ObjectID
}

/*
	Profile
*/

// Read
func (u *User) ReadProfile(mc *MongoConnection) (rest *application.Perfil) {

	bsonData, err := mc.FindOne(u.ID, "profile")
	if err != nil {
		fmt.Println(err)
	}

	err = bson.Unmarshal(bsonData, &rest)
	if err != nil {
		fmt.Println(err)
	}

	return rest
}

/*
	Deposits
*/
// Read
func (u *User) ReadDeposit(mc *MongoConnection) (rest *application.Deposits) {

	bsonData, err := mc.FindOne(u.ID, "deposits")

	if err != nil {
		fmt.Println(err)
	}

	err = bson.Unmarshal(bsonData, &rest)
	if err != nil {
		fmt.Println(err)
	}

	return rest
}

// Write
func (u *User) InsertDeposit(mc *MongoConnection, d application.Deposits) {

	doc := bson.D{
		{Key: "year_day", Value: d.YearDay},
		{Key: "deposits", Value: d.Deposits},
	}

	mc.InsetOne(doc, "deposits")

}
