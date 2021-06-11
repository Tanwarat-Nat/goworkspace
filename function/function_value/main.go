package main

import (
	"fmt"
	"math"
)

func add(a, b float64) float64 {
	return a + b
}

func sub(a, b float64) float64 {
	return a - b
}
func apply(a, b float64, op func(float64, float64) float64) (float64, error) { // op มีไทป์เป็น function signature
	if op == nil {
		return math.NaN(), fmt.Errorf("apply: nil operation")
	}
	return op(a, b), nil // <- เรียก invok func
}
func main() {
	a, _ := apply(1, 2, add)
	b, _ := apply(1, 2, sub)
	c, _ := apply(1, 2, nil)
	fmt.Println(a, b, c)
}

/*
เราจะไม่ใช้ function value แบบเดี่ยวๆ
program ที่มีความซับซ้อน มีความ complex มีการ design ที่ค่อนข้าง flexible
function value ค่อนข้างช่วยได้เยอะ ช่วยได้มากเลย
หรือใครที่เคยเขียน react แล้วเคยใช้ hiher component hiher order funtion
นี้คือทำงานคลายกันเลย
ก็คือ paste function value แล้วก็สามารถรีเทิร์นเป็น function value ออกมาได้ด้วย

*/
