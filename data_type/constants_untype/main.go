package main

import (
	"fmt"
	"reflect"
)

type Peice int

func main() {
	// Untype constants

	const person = 4
	tofee := Peice(5) / person
	costt := 2.01 / person
	fmt.Println(tofee, reflect.TypeOf(tofee))
	fmt.Println(costt, reflect.TypeOf(costt))
	// มันจะมี type ก็ต่อเมื่อ มันถูก evaluate ใน expression ที่มันอยู่

	const persons = 4
	toffee := 5 / persons
	cost := 2.01 / persons
	fmt.Println(toffee, reflect.TypeOf(toffee))
	fmt.Println(cost, reflect.TypeOf(cost))

	const persons2 = int(4)
	toffee2 := 5 / persons2
	cost2 := 2.01 / float64(persons2)
	fmt.Println(toffee2, reflect.TypeOf(toffee2))
	fmt.Println(cost2, reflect.TypeOf(cost2))

	// untype จะเป็น int, float ก็ได้ แต่จะไม่มีการระบุว่าเป็น 8, 16, 32
	// หาก person เป็น float เมื่อนำไปทำการ operation กับ value ตัวอื่น (toffee, cost)
	// type ของ value ตัวนั้น (toffee, cost)
	// จะแปรไปเป็น float64() ได้เท่านั้น
	// แต่ หาก person เป็น int type พวกนั้นก็จะคงเดิม ไม่เปลี่ยน
}
