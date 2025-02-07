package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"connectDatabase/model"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://admin:12345@cluster0.zqufr.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

const dbName = "netflix"
const colName = "watchList"

// Most Important
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

//MONGO dB Helpers- file

// insert one record
func insertOneMovie(movie model.Netflix) {
	// Example insert logic
	inserted, err := collection.InsertOne(context.TODO(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single movie: ", inserted.InsertedID)
}

// update one record

func updateOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)

	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err1 := collection.UpdateOne(context.Background(), filter, update)

	if err1 != nil {
		log.Fatal(err1)
	}

	fmt.Println("modified count : ", result.ModifiedCount)

}

// delete one record for now

func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)

	filter := bson.M{"_id": id}

	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Movie got deleted count : ", deleteCount)

}

// delete all ogf the movies
func deleteAllmovies() int64 {
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Numbr of movies delete: ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

// get all movies from the data base
func getAllMovies() []primitive.M {

	cursor, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}
	var movies []primitive.M

	for cursor.Next((context.Background())) {

		var movie bson.M
		err := cursor.Decode(&movie)

		if err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movie)
	}

	defer cursor.Close(context.Background())
	return movies
}

// Actual Controllers

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Conetent-Type", "application/x-www-form-urlearncode")

	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)

}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Conetent-Type", "application/x-www-form-urlearncode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)

}

func MarkedAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Conetent-Type", "application/x-www-form-urlearncode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Conetent-Type", "application/x-www-form-urlearncode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Conetent-Type", "application/x-www-form-urlearncode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllmovies()
	json.NewEncoder(w).Encode(count)
}
