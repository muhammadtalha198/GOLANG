package connectdb

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	connectionString = "mongodb+srv://admin:12345@cluster0.zqufr.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
	dbName           = "Todo_Database"
	colName          = "Todo_Collection"
)

var collection *mongo.Collection

func init() {

	//_client options
	clientOptions := options.Client().ApplyURI(connectionString)

	//connect to mongo db
	_client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongo db connected successfully.")

	err = _client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")

	collection = _client.Database(dbName).Collection(colName)

	//collection instance
	fmt.Println("connection instance is ready!")

}

// ConnectDB initializes the MongoDB client and returns a collection
func ConnectDB() *mongo.Collection {

	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")

	collection := client.Database(dbName).Collection(colName)
	return collection
}
