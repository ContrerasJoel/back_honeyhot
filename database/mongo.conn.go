package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetCollection(collection string) *mongo.Collection {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://root:1234qwas@atlascluster.2wzvams.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		panic(err.Error())
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		panic(err.Error())
	}

	return client.Database("honeyhot").Collection(collection)

}
