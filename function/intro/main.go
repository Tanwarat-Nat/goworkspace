package main

import (
	"fmt"
	"reflect"
)

/*func name(parameter-list) (result-list) {
	...
	...
}*/

func avg(a, b float64) (c float64) {
	c = (a + b) / 2
	a = 999
	return c
}

func avg2(a, b float64, c int, d, e, f int32) float64 {
	return (a + b) / 2
}

func avg3(a, b float64, c int, d, e, f int32) (float64, string, int) {
	return 4.0, "hello", 1
}

func avg4(_, _ float64, _ int, _, _, _ int32) (_ float64, _ string, _ int) {
	return 1.2, "string", 3

} // ถ้าเราไม่ใช้ก็สามารถ ignore -> _ มันได้

func avg5(a, b float64) float64 {
	c := (a + b) / 2
	a = 999
	return c
}

func main() {
	x := 1.0
	a := avg5(x, 3)
	fmt.Println(x)
	b := avg5(1, 2)
	fmt.Println(a, b)
	fmt.Println(reflect.TypeOf(avg)) // ดู function signature
	fmt.Println(reflect.TypeOf(avg2))
	fmt.Println(reflect.TypeOf(avg3))
	fmt.Println(avg3(4.0, 2.5, 3, 4, 5, 6)) // ต้องใส่ให้ครบห้ามละ
	fmt.Println(avg4(1.1, 1.2, 3, 4, 5, 6))
}

/*need -> proceuer มัน execute ซ้ำกัน ถ้าเรา abstract มันเป็น function
ความผิดพลาดมันก็จะน้อยลง
parameter-list -> จะมีหนึ่งตัว หลายตัว หรือไม่มีเลยก็ได้ type เหมือนกัน omit เหลือตัวเดียวได้
result-list  -> จะมีหนึ่งตัว หรือหลายตัวก็ได้
function signature -> parameter คืออะไร return คืออะไร ไม่ใช่ name แต่เป็น
type ของ patramerter ที่รับมาคืออะไร type ของ return คืออะไร
ต้องใส่ให้ครบ ไม่มี concept defualt parameter คือ set f = 0 ตอนบอก parameter
เพื่อที่จะได้ละตอนใส่ค่า
paste aggrement เป็น paste by value การเปลี่ยน parameter list ไม่กระทบต่อ
aggrement ที่เรา pate เข้ามา นอกจากว่าสิ่งที่เรา paste เข้ามามันเป็น reference type
(slice/map)มันอาจเปลี่ยน underlying data structure ตรงนั้นได้ ถึงแม้ว่าเราจะ paste slice
เข้ามาก็ตาม slice ตรงนั้นก็จะถูก copy คือ มันจะถูก copy ระหว่าง aggrement กับ
parameter ตรงนี้
*/
