package main

import (
	"fmt"
	"log"
	"net/http"

	route "github.com/muhammadtalha198/todaapp/Route"
)

func main() {

	fmt.Println("Go connect with mongo db")

	// Initialize the router
	r := route.InitializeRouter()

	fmt.Println("Server is starting .....")

	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
