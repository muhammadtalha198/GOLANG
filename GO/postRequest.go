package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("make a post request")
	PerformPostRequest()

}

func PerformPostRequest() {

	const myurl = "http://localhost:8000/post"

	//fake jason data

	requestBody := strings.NewReader(`
		{
			"coursename" : "Lets go with go lang",
			"price" : "0",
			"platform" : "learn code online"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))

}
