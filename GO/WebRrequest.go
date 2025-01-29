package main

import (
	"fmt"
	"io"
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
	defer response.Body.Close()

	dataBytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(dataBytes))
}
