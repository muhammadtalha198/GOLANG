package main

import (
	"connectDatabase/router"
	"fmt"
	"net/http"
)

//go get
//go get -u github.com/gorilla/mux
// go get go.mongodb.org/mongo-driver/v2/mongo

//

func main() {
	fmt.Println("Go connect with mongo db")
	r := router.Router()

	fmt.Println("Server is strting .....")

	http.ListenAndServe(":4000", r)
	fmt.Println("server is running on port 4000...")

}
