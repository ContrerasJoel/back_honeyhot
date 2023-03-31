package product

import (
	"context"
	"time"

	"github.com/ContrerasJoel/back_honeyhot/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = database.GetCollection("products")
var ctx = context.Background()

func Create(product Product) error {
	var err error
	_, err = collection.InsertOne(ctx, product)
	if err != nil {
		return err
	}
	return nil
}

func Read() (Products, error) {
	var products Products
	filter := bson.D{}
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var product Product
		err = cur.Decode(&product)
		if err != nil {
			return nil, err
		}

		products = append(products, &product)
	}
	return products, nil
}

func Update(product Product, productID string) error {
	var err error
	oid, _ := primitive.ObjectIDFromHex(productID)
	filter := bson.M{"_id": oid}
	update := bson.M{
		"$set": bson.M{
			"link":         product.Link,
			"title":        product.Title,
			"price":        product.Price,
			"descriptions": product.Descriptions,
			"update_at":    time.Now(),
		},
	}
	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func Delete(productID string) error {
	var err error
	var oid primitive.ObjectID
	oid, err = primitive.ObjectIDFromHex(productID)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": oid}
	_, err = collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
