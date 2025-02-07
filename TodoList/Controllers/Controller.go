package controller

import (
	"encoding/json"
	"net/http"

	helper "github.com/muhammadtalha198/todaapp/Helper"
	model "github.com/muhammadtalha198/todaapp/Model"
	connectdb "github.com/muhammadtalha198/todaapp/connectDb"
)

var todolist model.Todo

// Simple function to test that APIs are working
func healthCheck(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	res := model.Response{
		Msg:  "Health Check",
		Code: 200,
	}

	json.NewEncoder(w).Encode(res)

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

	json.NewEncoder(w).Encode("Todo created successfully")

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

	err = db.CreateTodo(todolist)
	if err != nil {
		json.NewEncoder(w).Encode("Failed to create todo")
		return
	}

	json.NewEncoder(w).Encode("Todo created successfully")

}

func UpdateTodoList(w http.ResponseWriter, r *http.Request) {
	// Implementation here
	
}

// func DeleteTodo(w http.ResponseWriter, r *http.Request) {
// 	// Implementation here
// }

// func DeleteTodoAll(w http.ResponseWriter, r *http.Request) {
// 	// Implementation here
// }

// func GetTodoById(w http.ResponseWriter, r *http.Request) {
// 	// Implementation here
// }

// func getAllTodosetAllTodos(w http.ResponseWriter, r *http.Request) {
// 	// Implementation here
// }
