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
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Creating JSON Data")

	fmt.Println("")

	fmt.Println("Encoding JSON Data")
	EncodeJson()

	fmt.Println("")

	fmt.Println("Decoding JSON Data")
	DecodeJson()
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

func DecodeJson() {
	jsonData := []byte(`
		[
			{
					"CourseName": "ReactJS",
					"CoursePrice": 1000,
					"Platform": "Udemy",
					"tags": [
							"web-dev",
							"frontend"
					]
			},
			{
					"CourseName": "Django",
					"CoursePrice": 2000,
					"tags": [
							"web-dev",
							"backend"
					]
			},
			{
					"CourseName": "Flutter",
					"CoursePrice": 1500,
					"Platform": "Udemy"
			}
		]
	`)

	var lcourses []course

	//Check if the JSON data is valid
	checkValid := json.Valid(jsonData)

	if checkValid {
		//Why pass reference? Because we want to modify the original slice
		json.Unmarshal(jsonData, &lcourses)

		fmt.Printf("%#v\n", lcourses)
	} else {
		fmt.Println("JSON data is not valid")
	}

	fmt.Println("")

	//Accessing the JSON data
	var myData []map[string]interface{}
	json.Unmarshal(jsonData, &myData)
	fmt.Printf("%#v\n", myData)

	fmt.Println("")

	for _, v := range myData {
		for key, value := range v {
			fmt.Printf("%s: %v\n", key, value)
		}
		fmt.Println("")
	}
}
