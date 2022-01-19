package database_mongodb

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/vicmanbrile/moneyGolang/profile"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Data struct {
	Perfil profile.Perfil
}

func (d *Data) Init() {
	ClientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_CONNECTION"))
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client, err := mongo.Connect(ctx, ClientOptions)
	if err != nil {
		log.Fatal(err)
	}

	collect := client.Database("moneyGolang").Collection("profile")

	var result bson.M

	objectId, err := primitive.ObjectIDFromHex("61e5dff48ccc2d6ee7ed063d")
	if err != nil {
		log.Println("Invalid id")
	}

	err = collect.FindOne(ctx, bson.D{{Key: "_id", Value: objectId}}).Decode(&result)
	if err != nil {
		panic(err)
	}

	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}

	PERFIL := &profile.Perfil{}
	err = json.Unmarshal(jsonData, &PERFIL)
	if err != nil {
		fmt.Printf("Error al convertir a JSON: %v", err)
	}

	d.Perfil = *PERFIL
}
