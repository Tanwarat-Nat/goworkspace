package main

import "fmt"

func main() {
	x := "hi-สวัสดี" //[]byte{0,1...,20}
	y := []byte{0x68, 0x69, 0x2d, 0xe0, 0xb8, 0xaa, 0xe0, 0xb8, 0xa7, 0xe0, 0xb8, 0xb1, 0xe0, 0xb8, 0xaa, 0xe0, 0xb8, 0x94, 0xe0, 0xb8, 0xb5}
	fmt.Println(x, len(x))
	fmt.Printf("deci = %d, hex = %x\n", 'h', 'h')
	fmt.Println("string(y)", string(y))

	for i := 0; i < len(x); i++ {
		fmt.Printf("0x%x, ", x[i])
	}

	fmt.Println("string(65)", string(65))
	fmt.Println("string(0x41)", string(0x41))
}

// สร้างแล้วเปลี่ยนไม่ได้
// %d = เลขฐาน 10
// %x = เลขฐาน 16
