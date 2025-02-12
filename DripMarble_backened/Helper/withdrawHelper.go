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

func (db *Database) UpdateWithdraw(withdrawID string, withdraw model.Withdraw) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Convert string ID to primitive.ObjectID
	id, err := primitive.ObjectIDFromHex(withdrawID)
	if err != nil {
		log.Println("Invalid withdraw ID:", err)
		return err
	}

	// Define filter and update
	filter := bson.M{"_id": id}
	update := bson.M{"$set": withdraw}

	// Perform the update
	result, err := db.WithdrawCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Println("Error updating withdraw:", err)
		return err
	}

	fmt.Println("Modified count for Withdraw:", result.ModifiedCount)
	return nil
}

func (db *Database) DeleteWithdraw(withdrawID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Convert string ID to primitive.ObjectID
	id, err := primitive.ObjectIDFromHex(withdrawID)
	if err != nil {
		log.Println("Invalid withdraw ID:", err)
		return err
	}

	// Define filter
	filter := bson.M{"_id": id}

	// Perform the delete operation
	result, err := db.WithdrawCollection.DeleteOne(ctx, filter)
	if err != nil {
		log.Println("Error deleting withdraw:", err)
		return err
	}

	fmt.Println("Deleted Withdraw count:", result.DeletedCount)
	return nil
}

func (db *Database) GetWithdrawByID(withdrawID string) (*model.Withdraw, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Convert string ID to primitive.ObjectID
	id, err := primitive.ObjectIDFromHex(withdrawID)
	if err != nil {
		log.Println("Invalid withdraw ID:", err)
		return nil, err
	}

	// Define filter
	filter := bson.M{"_id": id}

	// Perform the query
	var withdraw model.Withdraw
	err = db.WithdrawCollection.FindOne(ctx, filter).Decode(&withdraw)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("Withdraw not found:", err)
			return nil, fmt.Errorf("withdraw not found")
		}
		log.Println("Error finding withdraw:", err)
		return nil, err
	}

	return &withdraw, nil
}

func (db *Database) GetAllWithdraws() ([]model.Withdraw, error) {
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var withdraws []model.Withdraw
	filter := bson.D{{}}

	cursor, err := db.WithdrawCollection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var withdraw model.Withdraw
		err := cursor.Decode(&withdraw)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		withdraws = append(withdraws, withdraw)
	}

	return withdraws, nil
}
