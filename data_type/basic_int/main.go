package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.MinInt8, math.MaxInt8) // -128, 127
	fmt.Println(math.MaxUint8)              //รวม -+ = 255
	fmt.Println(math.MinInt16, math.MaxInt16)
}

// integers มี 2 แบบ
// signed (มีเครื่องหมาย - ) -> 8, 16, 32, 64 (นี่คือค่า bit)
// unsigned (ไม่มีเครื่องหมาย) -> 8, 16, 32, 64 (นี่คือค่า bit)
