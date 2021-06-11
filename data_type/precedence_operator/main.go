package main

import "fmt"

func main() {
	fmt.Println("2 + 3*5 =", 2+3*5)
	fmt.Println("5 % 2", 5%2)
	fmt.Println("-5 % 2", -5%2)
	fmt.Println("5 % -2", 5%-2)
	//เครื่องหมายของผลลัพธ์ ขึ้นอยู่กับตัวตั้ง

	fmt.Println("5 & 2", 5&2)
	// 101
	// 010
	// 0

	fmt.Println("5 & 5", 5&5)
	// 101
	// 101
	// 5

	fmt.Println("5 &^ 5", 5&^5)
	// 101
	// 010
	// 0

	fmt.Println("5 &^ 2", 5&^2)
	// 101
	// 101
	// 5

	fmt.Println(true && false || true)

}

// precedence operator
// 5   *   /   %   <<   >>   &   &^
// 4   +   -   |   ^
// 3   ==  !=  <   <=   >    >=
// 2   &&
// 1   ||
