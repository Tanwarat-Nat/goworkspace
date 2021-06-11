package main

import "fmt"

// name : "Filicity"
// age : 23

func main() {
	x := struct {
		name string
		age  int
	}{age: 23, name: "Filicity"}
	// สลับที่ไม่ได้ นอกจากจะระบุ filed name ไว้ด้วย และต้องระบุทุกอัน
	fmt.Println(x)
	fmt.Println(x.name)
	fmt.Println(x.age)

}

/*
arrays ถ้าประกาศ int ก็จะเป็น int ทั้งหมด
float ก็เป็น float หมด string ก็เป็น string
คือ ประกาศ type ไหนก็จะเป็นแบบเดียวกะันหมด

 struct ไม่จำเป็นต้องแบบเดียวกันทั้งหมด ลูกผสมได้
 จะเป็น int flost string  bool ผสมได้หลายอย่าง หลาย type
 หรือจะ contain struct ด้วยกันเองก็ได้ หรือ contain array ด้วยกันเองก็ได้
 หรือจะเป็น map ก้ได้ ได้หมด ไม่จำเป็นต้อง strick ว่าต้องเป็นอะไร

 A record(เรียก tuple or struct) - data structure ที่เอาหลายๆอย่างมารวมกัน
value ใหญ่ๆที่เรียกว่า struct มี value อื่นๆรวมอยู่ด้วย อย่าง
preson ก็ยังมี name surname age gender tel เป็นต้น

index by names
ใน map เป็น type เดียวทั้งหมด [string]/[int]
แต่ใน struct เป็นได้หลาย type ตามข้อมูล  value

*/
