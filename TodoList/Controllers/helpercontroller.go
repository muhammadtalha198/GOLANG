package controller

import (
	"context"
	"fmt"
	"log"
	"time"

	model "github.com/muhammadtalha198/todaapp/Model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Database struct to hold MongoDB collection reference
type Database struct {
	collection *mongo.Collection
}

// create a todo in the database
func (db *Database) CreateTodo(todolist model.Todo) {

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	inserted, err := db.collection.InsertOne(context.TODO(), todolist)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted with inserted Id: ", inserted.InsertedID)
}

// create a todo in the database with cudstom data
func (db *Database) CreateTodoCustom(todolist model.Todo) error {

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Prepare the data for insertion
	newTodo := model.Todo{
		Task:      todolist.Task,
		Completed: todolist.Completed,
		CreatedAt: time.Now(), // Set created time
		UpdatedAt: time.Now(), // Set updated time
	}

	inserted, err := db.collection.InsertOne(context.TODO(), newTodo)
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("Inserted with inserted Id: ", inserted.InsertedID)

	return nil
}

func (db *Database) UpdateTodoList(todoId string, todolist *model.Todo) error {

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id, err := primitive.ObjectIDFromHex(todoId)

	if err != nil {
		log.Fatal(err)
		return err
	}

	filterId := bson.M{"_id": id}

	//Single  value update
	// update := bson.M{"$set": bson.M{"completed": true}}

	update := bson.M{"$set": todolist}

	result, err1 := db.collection.UpdateOne(context.Background(), filterId, update)

	if err1 != nil {
		log.Fatal(err1)
		return err1
	}

	fmt.Println("modified count : ", result.ModifiedCount)

	return nil

}


