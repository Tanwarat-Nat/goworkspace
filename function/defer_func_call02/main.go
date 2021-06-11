package main

import "fmt"

func main() {
	x := 1
	defer fmt.Println("defer 1 x: d")
	defer fmt.Println("defer 2 x:", x+2)
	fmt.Println("x:", x)

	y := 1
	addY := func(a int) int {
		y = y + a
		return y
	}
	defer fmt.Println("defer 3 y:", addY(3)+5)
	fmt.Println("y:", y)
}

/*
Defer : Get to the last function call

defer จะทำการ evaluared(จัดการตัว expression ต่างๆ) ก่อน (1+2)
แล้วค่อย call ทีหลัง (จัดการแล้ว เอาค่า 3 ไปยัดที่ defer stack)
*/
