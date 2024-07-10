package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://dummy.restapiexample.com/api/v1/employees"

func main() {
	fmt.Println("Web Requests")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n", response)

	//Defer the closing of the response body because it is a good practice to close the response body after using it.
	defer response.Body.Close()

	// ioutil.ReadAll reads the response body and returns a byte slice
	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	// Convert the byte slice to a string
	content := string(databytes)

	fmt.Println(content)
}
