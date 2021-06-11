package main

import (
	"fmt"
)

func countDown(c int) {
	fmt.Println(c)
	if c == 0 {
		return
	}
	countDown(c - 1)
}

func main() {
	countDown(5)

}

/*func เรียกใช้ตัวมันเอง เรียก recusive หาก run ผลลัพธ์มันจะออกมาไม่สิ้นสุด ต้องใช้
ctrl + c เพื่อออกไป แต่มันจะไม่ใช่การ exit แบบ normal condition ที่จะทำให้มันหยุดตัวมันเอง
อาจจะสร้าง function countdown
ส่วนประกอบสำคัญเลยของ recursive function คือ terminate condition
หมายถึงว่า condition ไหนที่จะหยุดการ call function ตัวมันเอง
function ไหนที่จะ พอ condition นี้ apply ปุ๊บมันก็จะจบการทำงาน
แล้วไม่เรียกตัวมันเองซ้ำอีก
*/
