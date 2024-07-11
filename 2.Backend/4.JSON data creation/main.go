package main

import (
	"encoding/json"
	"fmt"
)

// The third parameter is the tag name which will be used in the JSON data
// `json:"-"` will not be included in the JSON data
// `json:"tgas,omitempty"` will be included in the JSON data if it is not nil
type course struct {
	Name     string `json:"CourseName"`
	Price    int    `json:"CoursePrice"`
	Platform string
	Password string   `json:"-"`
	Tags     []string `json:"tgas,omitempty"`
}

func main() {
	fmt.Println("Creating JSON Data")

	fmt.Println("")

	EncodeJson()
}

func EncodeJson() {
	lcourses := []course{
		{"ReactJS", 1000, "Udemy", "password", []string{"web-dev", "frontend"}},
		{"Django", 2000, "Udemy", "password", []string{"web-dev", "backend"}},
		{"Flutter", 1500, "Udemy", "password", nil},
	}

	//Package this data as JSON data
	finalJson, err := json.MarshalIndent(lcourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}
