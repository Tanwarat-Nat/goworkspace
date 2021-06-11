package main

import "fmt"

func main() {
	x := []int{1, 2, 3}
	fmt.Println(x[1])
	panic("for no reason") // panic => program stop suddenly
	fmt.Println(x[2])

}

/*
ในขณะที่โปรแกรมของเรา run อยู่ แล้วเกิดเหตุการณ์ไม่คาดคิดเกิดขึ้น
อย่างเช่น เรามี structure อันนึง แล้ว struct อันนั้นมันเป็น null
แล้วเราไปเรียก struct member จาก null pointer ตรงนั้น
มันก็จะขึ้น error เพราะมันไม่สามารถหาข้อมูลใน memory ได้

array ก็เช่นเดียวกัน สมมติว่าเรามี array ที่มีจำนวน element อยู่ที่ 3 ตัว
แล้วเราใช้ ascess index ไปที่ 100 มันก็จะเป็น index ที่ไม่อยู่ใน range ของมัน
มันก็จะ index out of range

concept คือ มัน error มันไม่สามารถทำต่อไปได้

panic ไม่จำเป็นต้องถูก generate มาจากโปรแกรมที่ทำงานผิดพลาดก็ได้
เราสามารถ simulate programe ให้อยู่ในสถานะ panic ได้
คือ panicking state ถ้าโปรแกรมอยู่ใน state ที่มี panic
โปรแกรมนั้นเรียกว่าอยู่ใน panicking state นั่นก็คือ สมมติว่า

พอ panic ปุ๊ป โปรแกรมจะจบการทำงานเลยทันที
อะไรที่อยู่ใต่ panic จะถูก skip ถูก ignor ไป ไม่ execute
โปรแกรมจะอยู่ในสถานะ panic

*/
