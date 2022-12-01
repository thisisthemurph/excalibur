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
	fmt.Println("Collection reference ready")
	return db, dbCtx
}
