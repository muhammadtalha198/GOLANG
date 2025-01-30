package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("make a get request")

	PerformgetRequest()
}

func PerformgetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("StatusCode is : ", response.StatusCode)
	fmt.Println("content length is : ", response.ContentLength)

	var responseString strings.Builder

	content, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	byteCount, err := responseString.Write(content)
	if err != nil {
		panic(err)
	}

	fmt.Println("byte count will be : ", byteCount)
	fmt.Println(responseString.String())

	// Optionally print the content
	// fmt.Println(responseString.String())
}
