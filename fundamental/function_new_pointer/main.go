package main

import "fmt"

func newIntPointer() *int {
	// initualize => zero value
	var x int
	return &x
}

func main() {

	fmt.Println(newIntPointer() == newIntPointer())

	// func newIntPointer() มี built-in function => new()
	fmt.Println(new(int) == new(int))

}
