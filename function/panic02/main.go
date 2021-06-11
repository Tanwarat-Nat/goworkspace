package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("fileName")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(f)

}

/*
คำถามพบบ่อย
ถ้าเรา check index เราเข้าถึง index อย่างเช่น x[20] เราเช็ค index ว่ามัน
น้อยกว่า range ของ x-1 รึเปล่า ก้ได้ ถ้ามันไม่เข้า condition จะ return error
ออกไป ซึ่งส่วนมากแล้วเขาจะทำแบบนี้ พวกไลบารี่ต่างๆเขาก็จะรีเทิร์น error ออกมา

open fileName: The system cannot find the file specified.
มันโอเคกว่า ตรงที่ os.Open จะ return error ออกมาแล้วให้เรา handle error
ตรงนี้เอง
ถ้าเกิดว่า os.Open หาไฟล์ไม่เจอแล้ว panic ออกมาเลย โปรแกรมเราก็จะถูก stop execute
จบการทำงานทุกอย่าง โปรแกรมมันไม่ควรจะเป็นแบบนั้น
ดังนั้นก็เป็นเหตุผลที่ว่ามีน้อยเคสมากที่เราจะใช้ panic ในโปรแกรมของเรา


*/
