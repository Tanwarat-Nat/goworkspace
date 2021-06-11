package main

import "fmt"

func main() {

	// if condition {

	// }

	if true {
		fmt.Println("I am in if")
	} else {
		fmt.Println("I am in else")
	}

	x := 5
	if x < 2 {
		fmt.Println("X in if")
	} else if x > 3 {
		fmt.Println("X in else")
	}

	y := 5
	if y <= 5 {
		fmt.Println("Y in if")
	} else {
		fmt.Println("Y in else")
	}

	// ใน && ถ้ามี false อันนึงก็จะ false ไปเลย
	z := 5
	if (z <= 5) && false {
		fmt.Println("Z in if")
	} else {
		fmt.Println("Z in else")
	}

	// ใน || ถ้ามี true อันนึงก็จะ true ไปเลย
	i := 5
	if (i <= 5) || false {
		fmt.Println("I in if")
	} else {
		fmt.Println("I in else")
	}

	j := 5
	if v := j + 1; (j <= 5) || false {
		fmt.Println("J in if", v)
	} else {
		fmt.Println("J in else", v)
	}

	k := 5
	if m := k + 1; (k <= 5) && false {
		fmt.Println("K in if", m)
	} else {
		fmt.Println("K in else", m)
	}
}
