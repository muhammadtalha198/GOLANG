package helper

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
	Collection *mongo.Collection
}

// Exported CreateTodo function
func (db *Database) CreateTodo(todolist model.Todo) error {
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	inserted, err := db.Collection.InsertOne(context.TODO(), todolist)
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("Inserted with inserted Id: ", inserted.InsertedID)
	return nil
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

	inserted, err := db.Collection.InsertOne(context.TODO(), newTodo)
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

	result, err1 := db.Collection.UpdateOne(context.Background(), filterId, update)

	if err1 != nil {
		log.Fatal(err1)
		return err1
	}

	fmt.Println("modified count : ", result.ModifiedCount)

	return nil
}

func (db *Database) DeleteTodo(todoId string) error {

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id, err := primitive.ObjectIDFromHex(todoId)

	if err != nil {
		log.Fatal(err)
		return err
	}

	filterId := bson.M{"_id": id}

	deleteCount, err1 := db.Collection.DeleteOne(context.Background(), filterId)

	if err1 != nil {
		log.Fatal(err1)
		return err1
	}

	fmt.Print("Movie got deleted count : ", deleteCount)

	return nil
}

func (db *Database) DeleteTodoAll() int64 {

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filterId := bson.D{{}}

	deleteResult, err1 := db.Collection.DeleteMany(context.Background(), filterId, nil)

	if err1 != nil {
		log.Fatal(err1)

	}

	fmt.Print("Movie got deleted count : ", deleteResult)

	return deleteResult.DeletedCount
}

func (db *Database) GetTodoById(todoId string) (*model.Todo, error) {

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id, err := primitive.ObjectIDFromHex(todoId)

	if err != nil {
		log.Fatal(err)
	}

	var todolist model.Todo

	filterId := bson.M{"_id": id}

	err = db.Collection.FindOne(context.Background(), filterId).Decode(&todolist)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &todolist, nil
}

func (db *Database) GetAllTodos() ([]model.Todo, error) {
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var todolists []model.Todo

	filterId := bson.D{{}}

	cursor, err := db.Collection.Find(context.TODO(), filterId)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for cursor.Next(context.Background()) {
		var todolist model.Todo
		err := cursor.Decode(&todolist)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		todolists = append(todolists, todolist)
	}
	defer cursor.Close(context.Background())

	return todolists, nil

}


