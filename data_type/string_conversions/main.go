package main

import "fmt"

func main() {
	// case 1: Integer to String
	ex1 := string(0x265e)
	fmt.Println("case 1: Integer -> String")
	for i := 0; i < len(ex1); i++ {
		fmt.Printf("%x ", ex1[i])
	}
	fmt.Println("\n\xe2\x99\x9e")
	fmt.Println("\xe2\x99\x9e" == ex1)

	// case 2: Slice of byte to String
	fmt.Println("\nCase 2: []byte -> string")
	ex2 := []byte{0xe2, 0x99, 0x9e}
	ex2String := string(ex2)
	fmt.Println(ex2String)

	// case 3: Slice of rune to String
	fmt.Println("\nCase 3: []rune -> string")
	ex3 := []rune{0x2654, 0x2655, 0x2656, 0x2657, 0x2658, 0x2659}
	fmt.Println(string(ex3))

	// case 4: String to Slice of byte
	fmt.Println("\nCase 4: String -> []byte")
	ex4 := []byte("hello♕")
	fmt.Println(ex4)

	// case 5: String to a Slice of rune
	fmt.Println("\nCase 5: String -> []rune")
	ex5 := []rune("hello♕")
	fmt.Println(ex5)

}

// ใน code ใส่เป็นค่า hex
// แต่เมื่อแสดงใน debug จะแสดงค่า deci

// 1 rune ไม่จำเป็ฌนต้องมี 1 byte
// แล้วแต่ character ที่เลือกมา
// อย่าง ส 1 rune มี 3 byte
