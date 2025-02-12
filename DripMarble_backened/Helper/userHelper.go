package helpers

import (
	"context"
	"fmt"
	"log"
	"time"

	model "github.com/muhammadtalha198/dripmarble/Models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateUserRegistered inserts a new user
func (db *Database) CreateUserRegistered(user model.UserRegistered) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user.ID = primitive.NewObjectID() // Assign a new ObjectID

	inserted, err := db.UserCollection.InsertOne(ctx, user)
	if err != nil {
		log.Println("Error inserting user:", err)
		return primitive.NilObjectID, err // Return empty ObjectID in case of error
	}

	fmt.Println("Inserted User with ID:", inserted.InsertedID)
	return inserted.InsertedID.(primitive.ObjectID), nil // Convert InsertedID to ObjectID
}

func (db *Database) UpdateUserRegistered(userID string, user model.UserRegistered) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Convert string ID to primitive.ObjectID
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Println("Invalid user ID:", err)
		return err
	}

	// Define filter and update
	filter := bson.M{"_id": id}
	update := bson.M{"$set": user}

	// Perform the update
	result, err := db.UserCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Println("Error updating user:", err)
		return err
	}

	fmt.Println("Modified count for UserRegistered:", result.ModifiedCount)
	return nil
}

func (db *Database) GetAllUsers() ([]model.UserRegistered, error) {
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var users []model.UserRegistered
	filter := bson.D{{}}

	cursor, err := db.UserCollection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var user model.UserRegistered
		err := cursor.Decode(&user)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
