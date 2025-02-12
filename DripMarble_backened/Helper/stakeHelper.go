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

// CreateStake inserts a new stake
func (db *Database) CreateStake(userID primitive.ObjectID, stake model.Stake) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stake.ID = primitive.NewObjectID() // Assign a new ObjectID
	stake.TimeStamp = time.Now()       // Set current timestamp as time.Time
	stake.UserID = userID

	inserted, err := db.StakeCollection.InsertOne(ctx, stake)
	if err != nil {
		log.Println("Error inserting stake:", err)
		return primitive.NilObjectID, err
	}

	fmt.Println("Inserted Stake with ID:", inserted.InsertedID)
	return stake.UserID, nil
}

func (db *Database) UpdateStake(stakeID string, stake model.Stake) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Convert string ID to primitive.ObjectID
	id, err := primitive.ObjectIDFromHex(stakeID)
	if err != nil {
		log.Println("Invalid stake ID:", err)
		return err
	}

	// Define filter and update
	filter := bson.M{"_id": id}
	update := bson.M{"$set": stake}

	// Perform the update
	result, err := db.StakeCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Println("Error updating stake:", err)
		return err
	}

	fmt.Println("Modified count for Stake:", result.ModifiedCount)
	return nil
}

func (db *Database) DeleteUserRegistered(userID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Convert string ID to primitive.ObjectID
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Println("Invalid user ID:", err)
		return err
	}

	// Define filter
	filter := bson.M{"_id": id}

	// Perform the delete operation
	result, err := db.UserCollection.DeleteOne(ctx, filter)
	if err != nil {
		log.Println("Error deleting user:", err)
		return err
	}

	fmt.Println("Deleted UserRegistered count:", result.DeletedCount)
	return nil
}

func (db *Database) DeleteStake(stakeID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Convert string ID to primitive.ObjectID
	id, err := primitive.ObjectIDFromHex(stakeID)
	if err != nil {
		log.Println("Invalid stake ID:", err)
		return err
	}

	// Define filter
	filter := bson.M{"_id": id}

	// Perform the delete operation
	result, err := db.StakeCollection.DeleteOne(ctx, filter)
	if err != nil {
		log.Println("Error deleting stake:", err)
		return err
	}

	fmt.Println("Deleted Stake count:", result.DeletedCount)
	return nil
}

func (db *Database) GetUserRegisteredByID(userID string) (*model.UserRegistered, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Convert string ID to primitive.ObjectID
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Println("Invalid user ID:", err)
		return nil, err
	}

	// Define filter
	filter := bson.M{"_id": id}

	// Perform the query
	var user model.UserRegistered
	err = db.UserCollection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("User not found:", err)
			return nil, fmt.Errorf("user not found")
		}
		log.Println("Error finding user:", err)
		return nil, err
	}

	return &user, nil
}

func (db *Database) GetStakeByID(stakeID string) (*model.Stake, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Convert string ID to primitive.ObjectID
	id, err := primitive.ObjectIDFromHex(stakeID)
	if err != nil {
		log.Println("Invalid stake ID:", err)
		return nil, err
	}

	// Define filter
	filter := bson.M{"_id": id}

	// Perform the query
	var stake model.Stake
	err = db.StakeCollection.FindOne(ctx, filter).Decode(&stake)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("Stake not found:", err)
			return nil, fmt.Errorf("stake not found")
		}
		log.Println("Error finding stake:", err)
		return nil, err
	}

	return &stake, nil
}

func (db *Database) GetAllStakes() ([]model.Stake, error) {
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var stakes []model.Stake
	filter := bson.D{{}}

	cursor, err := db.StakeCollection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var stake model.Stake
		err := cursor.Decode(&stake)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		stakes = append(stakes, stake)
	}

	return stakes, nil
}
