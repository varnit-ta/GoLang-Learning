package main

import (
	"fmt"
)

func variable_declaration() {
	//Declaring of variable
	var variable1 int = 10
	fmt.Println(variable1)

	//Declaration using :=
	variable2 := 1
	fmt.Println(variable2)

	//Declaration in block
	var (
		a int
		b int    = 1
		c string = "hello"
	)
	a = 0
	fmt.Println(a, "\n", b, "\n", c)

	//Go output %v -> value, %T -> type
	var (
		i string = "hello"
		j int    = 15
	)
	fmt.Printf("%v %T", j, i)
}
