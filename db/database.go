package db

import (
	"context"
	"os"

	"github.com/vicmanbrile/moneyGolang/profile"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetDataFindOne(id, coll string) ([]byte, error) {
	ClientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_CONNECTION"))
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client, err := mongo.Connect(ctx, ClientOptions)
	if err != nil {
		return nil, err
	}

	collect := client.Database("moneyGolang").Collection(coll)

	var result bson.M

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = collect.FindOne(ctx, bson.D{{Key: "_id", Value: objectId}}).Decode(&result)
	if err != nil {
		return nil, err
	}

	bsonData, err := bson.Marshal(result)
	if err != nil {
		return nil, err
	}

	return bsonData, nil
}

func GetDataProfile(id, coll string) (d *profile.Perfil, err error) {
	// Se extrae la información de GetDataFindOne() para comprovar si hay error o no se encontro el archivo
	profile, err := GetDataFindOne(id, coll)
	if err != nil {
		return nil, err
	}

	err = bson.Unmarshal(profile, &d)
	if err != nil {
		return nil, err
	}

	return
}
