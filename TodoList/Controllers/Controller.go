package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	helper "github.com/muhammadtalha198/todaapp/Helper"
	model "github.com/muhammadtalha198/todaapp/Model"
	connectdb "github.com/muhammadtalha198/todaapp/connectDb"
)

var todolist model.Todo

func HealthCheck(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	res := model.Response{
		Msg:  "Health Check",
		Code: 200,
	}

	json.NewEncoder(w).Encode(res)

}

func GetTodoById(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	// Check if the body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("please send some data.")
		return
	}

	// Extract todo ID from URL parameters
	params := mux.Vars(r)
	todoId, exists := params["id"]
	if !exists {
		http.Error(w, "Missing todo ID in URL", http.StatusBadRequest)
		return
	}

	// Create a Database instance with the initialized collection
	db := helper.Database{
		Collection: connectdb.ConnectDB(),
	}

	// Fetch the todo item from the database
	todo, err := db.GetTodoById(todoId)
	if err != nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	// Return the retrieved todo item as JSON
	json.NewEncoder(w).Encode(todo)

}

func GetAllTodos(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	// Check if the body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("please send some data.")
		return
	}

	// Create a Database instance with the initialized collection
	db := helper.Database{
		Collection: connectdb.ConnectDB(),
	}

	// Fetch all todo items from the database
	todos, err := db.GetAllTodos()
	if err != nil {
		http.Error(w, "Failed to fetch todos", http.StatusInternalServerError)
		return
	}

	// Return the list of todos as JSON
	json.NewEncoder(w).Encode(todos)

}

func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	// Check if the body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("please send some data.")
		return
	}

	var todolist model.Todo

	err := json.NewDecoder(r.Body).Decode(&todolist)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if todolist.IsEmpty() {
		json.NewEncoder(w).Encode("no data inside json.")
		return
	}

	// Create a Database instance with the initialized collection
	db := helper.Database{
		Collection: connectdb.ConnectDB(),
	}

	err = db.CreateTodo(todolist)
	if err != nil {
		json.NewEncoder(w).Encode("Failed to create todo")
		return
	}

	// json.NewEncoder(w).Encode("Todo created successfully")
	json.NewEncoder(w).Encode(todolist)

	// return todolist.Id;

	// // Include the ID in the response
	// response := map[string]interface{}{
	// 	"message": "Todo created successfully",
	// 	"id":      todolist.Id,
	// }

	// json.NewEncoder(w).Encode(response)
}

func CreateTodoCustom(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	// Check if the body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("please send some data.")
		return
	}

	var todolist model.Todo

	err := json.NewDecoder(r.Body).Decode(&todolist)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if todolist.IsEmpty() {
		json.NewEncoder(w).Encode("no data inside json.")
		return
	}

	// Create a Database instance with the initialized collection
	db := helper.Database{
		Collection: connectdb.ConnectDB(),
	}

	err = db.CreateTodoCustom(todolist)
	if err != nil {
		json.NewEncoder(w).Encode("Failed to create todo")
		return
	}

	json.NewEncoder(w).Encode("Todo created successfully")

}

func UpdateTodoList(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")

	// Check if the body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("please send some data.")
		return
	}

	params := mux.Vars(r)

	todoId, exists := params["id"]
	if !exists {
		http.Error(w, "Missing todo ID in URL", http.StatusBadRequest)
		return
	}

	var todolist model.Todo

	// Decode the request body into the todolist object
	err := json.NewDecoder(r.Body).Decode(&todolist)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create a Database instance with the initialized collection
	db := helper.Database{
		Collection: connectdb.ConnectDB(),
	}

	// Update the todo item in the database
	err = db.UpdateTodoList(todoId, &todolist)
	if err != nil {
		http.Error(w, "Failed to update todo", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("Todo updated successfully")
	// json.NewEncoder(w).Encode(params["id"])

}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Conetent-Type", "application/x-www-form-urlearncode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	// Check if the body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("please send some data.")
		return
	}

	params := mux.Vars(r)

	todoId, exists := params["id"]
	if !exists {
		http.Error(w, "Missing todo ID in URL", http.StatusBadRequest)
		return
	}

	// Create a Database instance with the initialized collection
	db := helper.Database{
		Collection: connectdb.ConnectDB(),
	}

	// Update the todo item in the database
	err := db.DeleteTodo(todoId)
	if err != nil {
		http.Error(w, "Failed to delete todo", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("Todo deleted successfully")
	// json.NewEncoder(w).Encode(params["id"])

}

func DeleteTodoAll(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Conetent-Type", "application/x-www-form-urlearncode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	// Check if the body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("please send some data.")
		return
	}

	db := helper.Database{
		Collection: connectdb.ConnectDB(),
	}

	// Update the todo item in the database
	count := db.DeleteTodoAll()

	// json.NewEncoder(w).Encode("Todo deleted successfully")
	json.NewEncoder(w).Encode(count)

}
