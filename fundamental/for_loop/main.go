package main

import (
	"fmt"
	"time"
)

func main() {

	//  for initialization; condition; post {
	// 	 	body
	//  }

	// การทำที่ง่ายสุดคือ omit ออกให้หมด -> infinite loop
	// มันจะ print ไปเรื่อยๆไม่สิ้นสุด -> Ctrl+C ออกไป

	// ถ้า omit คือ omit หมด
	// ถ้าใส่อันเดียว คือ condition จะเป็น boolean
	// ถ้าเป็น true body จะ execute ถ้า false body จะไม่ execute

	// for {
	// 	fmt.Println(time.Now().Clock())
	// }

	i := 1      // initialization
	for i < 5 { // condition
		fmt.Println(time.Now().Clock())
		i = i + 1 // i++ // post
	}

	// for initialization; condition; post
	for j := 1; j < 5; j++ {
		fmt.Println(j)
	}

	for k := 1; k < 5; k++ {
		fmt.Println(k)
		if k == 3 {
			fmt.Println("I am abount to terminate")
		}
	}

	// หยุดที่ 3 ไม่ต้อง print 4 ต่อ
	// break ออกจาก for loop ไม่ใช่ if loop
	for a := 1; a < 5; a++ {
		fmt.Println(a)
		if a == 3 {
			fmt.Println("I am abount to terminate")
			break
		}
	}

	// continune ไปเลยไม่ต้องสนใจ code ถัดไป
	// ไม่ต้อง print 2 และสิ่งใน {} ไปต่อเลย
	for b := 1; b < 5; b++ {
		if b == 2 {
			continue
		}
		fmt.Println(b)
	}
}
