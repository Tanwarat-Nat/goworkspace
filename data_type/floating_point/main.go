package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	fmt.Println(x, -x, 1/x, 1/-x, x/x)
	fmt.Println(math.IsNaN(x / x))
	fmt.Println(math.IsInf(1/x, 0))
	fmt.Println(math.IsInf(-1/x, 0))
	fmt.Println(math.IsInf(-1/x, 1))
	fmt.Println(math.IsInf(-1/x, -1))

	// ติดลบ ใน int (%) จะดูที่ตัวตั้ง
	// ติดลบ ใน float (/) จะดูที่ตัวหาร
	// 1/0.00000000001

}
