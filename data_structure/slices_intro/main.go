package main

import "fmt"

func main() {
	// สมมติมี ayyars อยู่ตัวนึง
	fruits := [5]string{
		"apple",
		"banana",
		"papaya",
		"orange",
		"mango",
	}
	fmt.Println("arrayrFruits =", fruits)

	// สร้าง slices จาข้อมูลของ arrays
	myFavor := fruits[1:4] // เอาตัวที่ 1-3 ไม่เอา 4
	fmt.Println("slicesMyfavor =", myFavor)

	/* internal structure of Slices
	data = [banana papaya orange] เป็น pointer ที่พ้อยไปที่ array นั้นๆ
	len = 3 จำนวนข้อมูลที่มีอยู่
	cap = 4 ค่าสูงสุดที่ให้เก็บข้อมูลได้
	*/

	yourFavor := myFavor                  // coppy แบบ value เลย
	yourFavor[0] = "guava"                // แต่เปลี่ยนส่วนที่เป็น pointer
	fmt.Println("yourFavor =", yourFavor) // data เลยเปลี่ยนตาม

}

// slices ใช้ arrays เป็นตัวเก็บข้อมูล
// arrays เองไม่รู้หรอกว่ามี slices ไหนบ้างมาจับกับข้อมูลของมัน
