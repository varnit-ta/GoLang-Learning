package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for course - file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"-"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// Fake DB
var courses []Course

// Main function - file
func main() {
	fmt.Println("Starting the application...")

	// Seeding
	courses = append(courses, Course{CourseId: "1", CourseName: "Java", CoursePrice: 100, Author: &Author{Fullname: "John Doe", Website: "www.johndoe.com"}})
	courses = append(courses, Course{CourseId: "2", CourseName: "Python", CoursePrice: 200, Author: &Author{Fullname: "Jane Doe", Website: "www.janedoe.com"}})

	// Routing
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/api/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/api/courses/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/api/courses", createOneCourse).Methods("POST")
	r.HandleFunc("/api/courses/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/api/courses/{id}", deleteOneCourse).Methods("DELETE")

	// Listen and serve
	log.Fatal(http.ListenAndServe(":4000", r))
}

// Middleware, helper - file
func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

// Controllers - file
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to the home page</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")

	//Set the header syntax: w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Type", "application/json")

	//Encode the courses and write it to the response sytax
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")

	w.Header().Set("Content-Type", "application/json")

	//Grab id from request
	params := mux.Vars(r)

	//Iterate over the courses and find the course with the id
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No Course found with given ID")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create a course")

	w.Header().Set("Content-Type", "application/json")

	//If body is empty, return
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send a request body")
	}

	//What about - {}
	var course Course

	//Decode the body and store it in course
	json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Please send a valid request body")
		return
	}

	//Generate uniques id and then convert it into string
	rand.Seed(time.Now().UnixNano())

	//"strconv.Itoa" converts int to string, "Itoa" stands for integer to ascii
	course.CourseId = strconv.Itoa(rand.Intn(1000000))

	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update a course")
	w.Header().Set("Content-Type", "application/json")

	//Grab id from request
	params := mux.Vars(r)

	// Loop through the courses and find the course with the id
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)

			var course Course
			json.NewDecoder(r.Body).Decode(&course)

			course.CourseId = params["id"]

			courses = append(courses, course)

			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No Course found with given ID")
	return
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete a course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course deleted successfully")
			return
		}
	}

	json.NewEncoder(w).Encode("No Course found with given ID")
	return
}
