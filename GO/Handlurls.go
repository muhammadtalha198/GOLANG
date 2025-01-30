package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://loc.dev:3000/learn?coursename=reactjs&paymentid=ghhj456fgh"

func main() {
	fmt.Println("Welcome to handle Urls")
	fmt.Println(myurl)

	// Parsing
	result, err := url.Parse(myurl)

	if err != nil {
		panic(err)
	}

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Println("The query params are:", qparams)

	fmt.Println("The query params are:", qparams["coursename"])

	for key, values := range qparams {
		fmt.Printf("Key: %s, Values: %v\n", key, values)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "looc.dev",
		Path:    "/tutcss",
		RawPath: "user=muhammad",
	}

	realUrl := partsOfUrl.String()
	fmt.Println(realUrl)
}
