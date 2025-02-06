package connectdb

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://admin:12345@cluster0.zqufr.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

const dbName = "Todo_Database"
const colName = "Todo_Collection"

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

	collection = _client.Database(dbName).Collection(colName)

	//collection instance
	fmt.Println("connection instance is ready!")

}
