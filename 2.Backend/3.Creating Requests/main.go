package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Creating Requests")

	fmt.Println("")

	fmt.Println("Performing GET Request")
	PerformGetRequest()

	fmt.Println("")

	fmt.Println("Performing POST Request with JSON")
	PerformPostJsonRequest()

	fmt.Println("")

	fmt.Println("Performing POST Request with Form")
	PerformPostFormRequest()
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

func PerformPostJsonRequest() {
	const myurl = "http://localhost:8000/post"

	// Creating a JSON payload
	requestBody := strings.NewReader(`
		{
			"courseName": "Golang",
			"price": 100,
			"platform": "Udemy"
		}
	`)

	// Sending the POST request (Syntax: http.Post(url, contentType, body)
	resp, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Println("Response Status: ", resp.Status)

	// Reading the response
	databytes, err := io.ReadAll(resp.Body)
	content := string(databytes)

	fmt.Println("Response Content: ", content)
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	// Creating a form payload
	data := url.Values{}
	data.Add("name", "John Doe")
	data.Add("age", "25")

	fmt.Println("Form Data: ", data)

	// Sending the POST request (Syntax: http.PostForm(url, data)
	resp, err := http.PostForm(myurl, data)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Println("Response Status: ", resp.Status)

	// Reading the response
	databytes, err := io.ReadAll(resp.Body)
	content := string(databytes)

	fmt.Println("Response Content: ", content)
}
