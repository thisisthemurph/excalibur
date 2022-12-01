package repository

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	databaseName     = "demo"
	collectionName   = "datatemplates"
	connectionString = "mongodb+srv://mike:NImEhZKCwLTbgt6q@development-cluster.vqxd1.mongodb.net/?retryWrites=true&w=majority"
)

// initiates the connection with the mongo database
func connect() (*mongo.Database, context.Context) {
	clientOptions := options.Client().ApplyURI(connectionString)
	dbCtx := context.TODO()
	client, err := mongo.Connect(dbCtx, clientOptions)

	db := client.Database(databaseName)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to mongo database successfully")

	// collection = client.Database(databaseName).Collection(collectionName)
	fmt.Println("Collection reference ready")
	return db, dbCtx
}

// func connect() (*mongo.Database, context.Context) {
// 	fmt.Println("Connecting to the database")

// 	// var collection *mongo.Collection

// 	clientOptions := options.Client().ApplyURI(connectionString)
// 	client, err := mongo.NewClient(clientOptions)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	ctx := context.Background()
// 	err = client.Connect(ctx)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// defer client.Disconnect(ctx)
// 	demoDB := client.Database(databaseName)

// 	// Testing - will fail if the connection exists
// 	// err = demoDB.CreateCollection(ctx, "datatemplates")
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }

// 	// dataTemplateCollection := demoDB.Collection("datatemplates")

// 	// Testing - removes the collection at the end of every run
// 	// defer dataTemplateCollection.Drop(ctx)

// 	// result, err := dataTemplateCollection.InsertOne(ctx, bson.D{
// 	// 	{Key: "name", Value: "Mocha"},
// 	// 	{Key: "breed", Value: "Turkish Van"},
// 	// })
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }

// 	// fmt.Println("Result: ", result)

// 	return demoDB, ctx
// }
