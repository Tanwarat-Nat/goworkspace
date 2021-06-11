package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	// อันนี้ทำให้เป็น anonymous filed ฟีลด์ที่ไม่มีชื่อ ก้คือลบชื่อออก
	Person
	designation string
}

func main() {
	filicity := Employee{
		Person{"Filicity", 23},
		"Programmer",
		// อันนี้ตาม anonym และการระบุ filed name ที่มีก็ต้องมีทั้งหมด ถ้าไม่มีก็ต้องไม่มีเลย
	}
	fmt.Printf("%+v\n", filicity)
	fmt.Println(filicity.name)
	fmt.Println(filicity.age)

	// ก็ไม่เชิง anonymous ซะทีเดียว มันก็มีชื่อของมันอยู่คือ Person.name
}
