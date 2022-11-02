package db

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConnection struct {
	Client           *mongo.Client
	ContexWithCancel context.Context
	ContexCancel     context.CancelFunc
}

func EstablishingConnection() (mc *MongoConnection) {

	mc = &MongoConnection{}

	ClientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_CONNECTION"))

	mc.ContexWithCancel, mc.ContexCancel = context.WithCancel(context.Background())

	mc.Client, _ = mongo.Connect(mc.ContexWithCancel, ClientOptions)

	return

}

func (mc *MongoConnection) CancelConection() {
	mc.ContexCancel()
}

func (mc *MongoConnection) FindOne(key primitive.ObjectID, collection string) ([]byte, error) {

	collect := mc.Client.Database("moneyGolang").Collection(collection)

	var result bson.M

	err := collect.FindOne(mc.ContexWithCancel, bson.D{{Key: "_id", Value: key}}).Decode(&result)
	if err != nil {
		return nil, err
	}

	bsonData, err := bson.Marshal(result)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(bsonData))

	return bsonData, nil
}
