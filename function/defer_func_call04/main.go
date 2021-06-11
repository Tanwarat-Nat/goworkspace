package main

import "fmt"

func add(a, b int) (c int) {
	defer func() {
		fmt.Printf("c is %d\n", c)
	}()
	return a + b
}

func main() {
	fmt.Println(add(2, 3))
}

/*return จะถูก execute ก่อน defer*/
