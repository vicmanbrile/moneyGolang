package db

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetData() []byte {
	ClientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_CONNECTION"))
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client, err := mongo.Connect(ctx, ClientOptions)
	if err != nil {
		log.Fatal(err)
	}

	collect := client.Database("moneyGolang").Collection("profile")

	var result bson.M

	objectId, err := primitive.ObjectIDFromHex("6215c7dc38821f527b019d3e")
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

	return jsonData
}
