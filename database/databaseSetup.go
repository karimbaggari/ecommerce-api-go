package database

import (
	"context"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBSet() *mongo.client{
	client,err - mongo.Newclient(options.Client().ApplyUrl("mongodb://localhost:27017"))

	if err!= nil{
        log.Fatal(err)
    }

	ctx, cancel = context.WithTimeOut(context.background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err!= nil{
		log.Fatal(err)
	}
	client.Ping(context.TODO(), nil)
	if err!= nil{
		log.Println("failed to connect to mongodb")
		return nil
	}
	fmt.Println("client connected")
	return client

}

var client *mongo.client = DBSet()


func UserData(client *mongo.Client, collectionName string) *mongo.client {
	var collection *mongo.Collection = client.database("Ecommerce").Collection(collectionName)
	return collection
}

func ProductData(client *mongo.Client, collectionName string) *mongo.client {
	var productCollection *mongo.Collection = client.database("Ecommerce").Collection(collectionName)
	return productCollection
}