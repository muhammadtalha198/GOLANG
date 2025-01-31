package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("make s server in go.")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serverHome).Methods("GET")

	// http.ListenAndServe(":4000",r)
	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	fmt.Println("Hey there mod userers")
}

func serverHome(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("<h1> Welcome to go lang series </h1>"))

}
