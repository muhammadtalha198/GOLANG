package helpers

import (
	"context"
	"fmt"
	"log"
	"time"

	model "github.com/muhammadtalha198/dripmarble/Models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (db *Database) CreateContract(contract model.Contract) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	contract.ID = primitive.NewObjectID() // Assign a new ObjectID

	inserted, err := db.UserCollection.InsertOne(ctx, contract)
	if err != nil {
		log.Println("Error inserting user:", err)
		return primitive.NilObjectID, err // Return empty ObjectID in case of error
	}

	fmt.Println("Inserted User with ID:", inserted.InsertedID)
	return inserted.InsertedID.(primitive.ObjectID), nil // Convert InsertedID to ObjectID
}

func (db *Database) GetContractByID(contractID string) (*model.Contract, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Convert string ID to primitive.ObjectID
	id, err := primitive.ObjectIDFromHex(contractID)
	if err != nil {
		log.Println("Invalid contract ID:", err)
		return nil, err
	}

	// Define filter
	filter := bson.M{"_id": id}

	// Perform the query
	var contract model.Contract
	err = db.ContractCollection.FindOne(ctx, filter).Decode(&contract)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("Contract not found:", err)
			return nil, fmt.Errorf("contract not found")
		}
		log.Println("Error finding contract:", err)
		return nil, err
	}

	return &contract, nil
}

func (db *Database) updateContract(contractId string, contract model.Contract) error {

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id, err := primitive.ObjectIDFromHex(contractId)

	if err != nil {
		log.Fatal(err)
		return err
	}

	filterId := bson.M{"_id": id}

	update := bson.M{"$set": contract}

	result, err1 := db.ContractCollection.UpdateOne(context.Background(), filterId, update)

	if err1 != nil {
		log.Fatal(err1)
		return err1
	}

	fmt.Println("modified count : ", result.ModifiedCount)

	return nil

}

func (db *Database) GetAllContracts() ([]model.Contract, error) {
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var contracts []model.Contract
	filter := bson.D{{}}

	cursor, err := db.ContractCollection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var contract model.Contract
		err := cursor.Decode(&contract)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		contracts = append(contracts, contract)
	}

	return contracts, nil
}

func (db *Database) DeleteContract(contractID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Convert string ID to primitive.ObjectID
	id, err := primitive.ObjectIDFromHex(contractID)
	if err != nil {
		log.Println("Invalid contract ID:", err)
		return err
	}

	// Define filter
	filter := bson.M{"_id": id}

	// Perform the delete operation
	result, err := db.ContractCollection.DeleteOne(ctx, filter)
	if err != nil {
		log.Println("Error deleting contract:", err)
		return err
	}

	fmt.Println("Deleted Contract count:", result.DeletedCount)
	return nil
}

func (db *Database) DeleteAllContracts() int64 {
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.D{{}}

	deleteResult, err := db.ContractCollection.DeleteMany(context.Background(), filter, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted Contracts Count:", deleteResult.DeletedCount)

	return deleteResult.DeletedCount
}
