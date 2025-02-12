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
	dbName           = "Drip_Marble"
)

var (
	client *mongo.Client
	db     *mongo.Database
)

// InitDB initializes the MongoDB client and collections
func InitDB() error {
	clientOptions := options.Client().ApplyURI(connectionString)

	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	// Ping the database to verify the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return fmt.Errorf("failed to ping MongoDB: %v", err)
	}

	fmt.Println("Connected to MongoDB!")

	db = client.Database(dbName)
	return nil
}

// CreateCollection returns a reference to a MongoDB collection with the given name
func CreateCollection(collectionName string) *mongo.Collection {
	return db.Collection(collectionName)
}

// ConnectDB connects to MongoDB and returns a collection
func ConnectDB(colName string) *mongo.Collection {
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

// CloseDB closes the MongoDB connection
func CloseDB() {
	if client != nil {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatalf("Failed to disconnect from MongoDB: %v", err)
		}
		fmt.Println("Disconnected from MongoDB!")
	}
}
