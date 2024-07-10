package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://www.google.com/search?q=golang&rlz=1C1GCEU_enIN829IN829&oq=golang&aqs=chrome..69i57j0l5.1161j0j7&sourceid=chrome&ie=UTF-8"

func main() {
	fmt.Println("Handling URL")
	fmt.Println(myurl)

	// Parsing URL
	result, _ := url.Parse(myurl)

	// Printing the result
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())
	fmt.Println(result.Fragment)

	// Getting the Query Params
	qparams := result.Query()
	fmt.Printf("Query Params: %v\n", qparams)
	fmt.Printf("The type of Query Params: %T\n", qparams)
	fmt.Printf("The value of Query Params: %v\n", qparams["q"])

	// Looping through the Query Params
	for key, value := range qparams {
		fmt.Printf("Key: %v, Value: %v\n", key, value)
	}

	// Creating a new URL (always pass reference to the URL)
	partsOfURL := &url.URL{
		Scheme:   "https",
		Host:     "www.google.com",
		Path:     "/search",
		RawQuery: "q=golang",
		RawPath:  "search?q=golang",
	}

	// Printing the new URL
	fmt.Println(partsOfURL.String())
}
