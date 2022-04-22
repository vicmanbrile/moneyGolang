package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/vicmanbrile/moneyGolang/profile"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetDataFindOne(id, coll string) []byte {
	ClientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_CONNECTION"))
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client, err := mongo.Connect(ctx, ClientOptions)
	if err != nil {
		log.Fatal(err)
	}

	collect := client.Database("moneyGolang").Collection(coll)

	var result bson.M

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid id")
	}

	err = collect.FindOne(ctx, bson.D{{Key: "_id", Value: objectId}}).Decode(&result)
	if err != nil {
		panic(err)
	}

	bsonData, err := bson.Marshal(result)
	if err != nil {
		panic(err)
	}

	return bsonData
}

func GetDataProfile(id, coll string) (d *profile.Perfil) {
	err := bson.Unmarshal(GetDataFindOne(id, coll), &d)

	if err != nil {
		fmt.Printf("Error al convertir a BSON: %v", err)
	}

	return
}
