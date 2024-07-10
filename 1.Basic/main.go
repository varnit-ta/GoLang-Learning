package main

import (
	"fmt"
	"math"
	"strings"
)

// Functions
func sayGreeting(n string) {
	fmt.Printf("Good morning %v \n", n)
}

func sayBye(n string) {
	fmt.Printf("Goodbye %v \n", n)
}

// Function with a Slice and Fucntion as a paramater
func cycleName(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

// Return type functions
func circleArea(r float64) float64 {
	return math.Pi * r * r
}

// Multiple vlauie return functions
func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ") //Returns a slice

	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}

func main() {
	// sayGreeting("John")
	// sayBye("John")

	// cycleName([]string{"cloud", "tifa"}, sayGreeting)

	// a1 := circleArea(12)
	// a2 := circleArea(10.5)
	// fmt.Println(a1, a2)

	// fn, sn := getInitials("tiffa lockhart")
	// fmt.Println(fn, sn)

	//Go maps
	menu := map[string]float64{
		"soup":               4.99,
		"pie":                7.99,
		"chocolate icecream": 10,
	}
	fmt.Println(menu)
	fmt.Println(menu["pie"])
	for key, value := range menu {
		fmt.Println(key, value)
	}
}
