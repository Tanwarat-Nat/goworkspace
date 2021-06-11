package main

import "fmt"

func fp() *int {
	x := 4
	return &x
}

func main() {
	x := fp()
	fmt.Printf("%T, %d\n", x, *x)
	fmt.Println(x == fp())

}

// ทุกครั้งที่มีการเรียก fp() local จะถูกสสร้างใหม่ทุกครั้ง
// และ return address ใหม่มาให้
// สร้าง stack ใน memory และ return local ใหม่มาให้
