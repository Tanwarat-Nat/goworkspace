package main

import (
	"fmt"
)

func main() {
	x := 2
	// const y = 0
	const z = 5 - 6
	//fmt.Println(x / y) //  รู้ว่าเป็น const แสดง error ที่
	//fmt.Println(x / 0) // compile time
	fmt.Println(x / z)

}

// ค่าที่ประกาศแล้วไม่สามารถเปลี่ยนแปลงได้ตลอดที่โปรแกรม execution
// การ reassign ไปให้ constant จะเกิด error ทันที
// จะไมาสามารถ complie ได้เลย
// Complier knows it values :
// Complier จะรู้จัก values
// Evaluation at compile time :
// Evaluation จะทำที่ compile time ไม่ใช่ run time
// คือ ก่อนจะเป็น exe ไฟล์ expression ที่เกี่ยวข้องกับ constant
// จะถูก evaluate มาเรียบร้อยแล้ว
// Declare constants as agroup :
// Iota
// Untype constants
