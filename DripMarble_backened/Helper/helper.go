package helpers

import (
	"context"
	"fmt"
	"log"
	"time"

	model "github.com/muhammadtalha198/dripmarble/Models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Database struct to hold MongoDB collection references
type Database struct {
	UserCollection     *mongo.Collection
	StakeCollection    *mongo.Collection
	WithdrawCollection *mongo.Collection
	ContractCollection *mongo.Collection
}

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

// CreateWithdraw inserts a new withdrawal
func (db *Database) CreateWithdraw(userID primitive.ObjectID, withdraw model.Withdraw) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	withdraw.ID = primitive.NewObjectID() // Assign a new ObjectID
	withdraw.TimeStamp = time.Now()       // Convert time.Time to primitive.DateTime
	withdraw.UserID = userID

	inserted, err := db.WithdrawCollection.InsertOne(ctx, withdraw)
	if err != nil {
		log.Println("Error inserting withdrawal:", err)
		return primitive.NilObjectID, err
	}

	fmt.Println("Inserted Withdraw with ID:", inserted.InsertedID)
	return withdraw.UserID, err
}
