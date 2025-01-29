package main

import (
	"fmt"
	"net/http"
)

const url = "https://loc.dev"

func main() {
	fmt.Println("Local web request")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of typ: %T\n", response)
	response.Body.Close()
	

}
