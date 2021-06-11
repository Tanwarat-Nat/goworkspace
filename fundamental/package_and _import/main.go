package main

import (
	"fmt"
	"weight" // เป็น path ของไฟล์
	// weight "weightx" - > package "Path"
	// เพื่อความง่ายจะตั้งชื่อ package กับ Path ตัวสุดท้ายให้เหมือนกัน
)

func main() {
	k := weight.KG(1.5) // เป็นชื่อ package
	fmt.Println(weight.KgToLb(k))
}
