package main

import "fmt"

func main() {
	x := int8(127)
	y := int8(1)
	fmt.Println("x + y =", x+y)

	a := 127
	b := 2
	fmt.Println("a + b =", int8(a+b))
	fmt.Printf("129 = %b\n", 129)

	c := 3.9999
	fmt.Println("c =", int(c))

}
