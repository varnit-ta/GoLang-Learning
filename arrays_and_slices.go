package main

import (
	"fmt"
)

// Declaring arrays
func arrays_and_slices() {
	arr1 := [5]int{} //not initialized

	var array = [5]int{1, 2, 3, 4, 5} //initialized

	var array1 = [...]int{1, 2, 3, 4, 5} //lenght not set
	array3 := [...]int{1, 2, 3, 4, 5}

	fmt.Println(arr1)
	fmt.Println(array)
	fmt.Println(array1)
	fmt.Println(array3)

	//initializing only specific elements of an array
	arr2 := [5]int{1: 10, 3: 30}
	fmt.Println(arr2)

	//Finding length of an array
	fmt.Println(len(arr1))

	//Declaring slices
	myslice1 := []int{} //not initialized

	myslice2 := []int{1, 2, 3, 4, 5} //initialized

	myslice3 := make([]int, 5) //length set to 5

	myslice4 := make([]int, 5, 10) //length set to 5, capacity set to 10

	fmt.Println(myslice1)
	fmt.Println(myslice2)
	fmt.Println(myslice3)
	fmt.Println(myslice4)
}
