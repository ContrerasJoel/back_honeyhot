package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetCollection(collection string) *mongo.Collection {
	mongoUser := os.Getenv("MONGOUSER")
	mongoPassword := os.Getenv("MONGOPASSWORD")
	uri := ""
	if mongoUser == "" {
		uri = "mongodb://localhost:27017"
	} else {
		uri = fmt.Sprintf("mongodb+srv://%s:%s@atlascluster.2wzvams.mongodb.net/?retryWrites=true&w=majority", mongoUser, mongoPassword)
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(uri)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		panic(err.Error())
	}

	return client.Database("honeyhot").Collection(collection)

}
