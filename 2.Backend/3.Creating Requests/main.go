package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Creating Requests")

	fmt.Println("Performing GET Request")
	PerformGetRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	resp, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Println("Response Status: ", resp.Status)
	fmt.Println("Content length: ", resp.ContentLength)

	// Reading the response
	databytes, err := io.ReadAll(resp.Body)
	content := string(databytes)

	fmt.Println("Response Content: ", content)

	// Writing the response using strings.Builder
	var responseString strings.Builder

	byteCount, _ := responseString.Write(databytes)
	responseContent := responseString.String()

	fmt.Println("Byte Count: ", byteCount)
	fmt.Println("Response Content: ", responseContent)
}
