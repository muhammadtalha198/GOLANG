package controller

import (
	"encoding/json"
	"log"
	"net/http"
)

// Define the Response struct
type Response struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}

// Simple function to test that APIs are working
func healthCheck(w http.ResponseWriter, r *http.Request) {
	res := Response{
		Msg:  "Health Check",
		Code: 200,
	}

	jsonStr, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Code)
	w.Write(jsonStr)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	// Implementation here
}

func CreateTodoCustom(w http.ResponseWriter, r *http.Request) {
	// Implementation here
}

func UpdateTodoList(w http.ResponseWriter, r *http.Request) {
	// Implementation here
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	// Implementation here
}

func DeleteTodoAll(w http.ResponseWriter, r *http.Request) {
	// Implementation here
}

func GetTodoById(w http.ResponseWriter, r *http.Request) {
	// Implementation here
}

func getAllTodosetAllTodos(w http.ResponseWriter, r *http.Request) {
	// Implementation here
}
