package main

import "fmt"

func main() {
	x := -5
	fmt.Printf("%b\n", x)
	fmt.Println(bitInt(int8(x)))
}

// 1111 1011
// 0000 0001
//         1-1 -> True
func bitInt(x int8) [8]byte {
	var result [8]byte // 1111 1011 (-5)
	for i := 0; i < len(result); i++ {
		result[7-i] = byte(x & 1)
		x = x >> 1
		// 1111 1011
		// 01111 101 ใส่ศูนย์ข้างหน้า แล้วตัดส่วนเกินด้านท้ายออก
		// จนกว่าตรงนี้จะเป็นศูนย์ทั้งหมด
	}
	return result
}
