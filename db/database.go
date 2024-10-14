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
	MongoClient      *mongo.Client
	ContexWithCancel context.Context
	FuncCancel       context.CancelFunc
}

func NewMongoConnection() *MongoConnection {
	mc := &MongoConnection{}
	var err error

	ClientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_CONNECTION"))

	mc.ContexWithCancel, mc.FuncCancel = context.WithCancel(context.Background())

	mc.MongoClient, err = mongo.Connect(mc.ContexWithCancel, ClientOptions)
	if err != nil {
		fmt.Println(err)
	}

	return mc
}

func (mc *MongoConnection) CancelConection() {
	mc.FuncCancel()
}

func (mc *MongoConnection) FindOne(key primitive.ObjectID, collection string) ([]byte, error) {

	dbs := mc.MongoClient.Database("moneyGolang")

	collect := dbs.Collection(collection)

	var result bson.M

	err := collect.FindOne(mc.ContexWithCancel, bson.D{{Key: "_id", Value: key}}).Decode(&result)
	if err != nil {
		return nil, err
	}

	bsonData, err := bson.Marshal(result)
	if err != nil {
		return nil, err
	}

	return bsonData, nil
}

func (mc *MongoConnection) InsetOne(document bson.D, collection string) {
	coll := mc.MongoClient.Database("moneyGolang").Collection(collection)

	result, err := coll.InsertOne(context.TODO(), document)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
}
