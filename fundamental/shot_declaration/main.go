package main

import "fmt"

// shot declaration cannot decla outsind funtion body
// only normal declaration can (var name type = expression)

func main() {

	// name := expression

	x := 1
	y, z := 2, 3

	fmt.Println(x)
	fmt.Println(y, z)

	// multiple(name) := multiple(expression)

	a, b := 1, 2
	b, c := a, b

	fmt.Println(a, b, c)
}
