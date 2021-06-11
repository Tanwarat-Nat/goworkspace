package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func (p *Point) Length() float64 {
	return math.Hypot((*p).X, (*p).Y)
}

func (p *Point) MoveXTo(newX float64) { // หน้าตาแบบนี้ คือ method
	(*p).X = newX
}

func (p Point) MoveYToByValue(newY float64) { // method เรียกใช้ function
	MoveYTo(&p, newY)
}

func (p *Point) MoveYTo(newY float64) {
	MoveYTo(p, newY)
}

func MoveYTo(p *Point, newY float64) { // หน้าตาแบบนี้ คือ function
	(*p).Y = newY
}

func main() {
	p := Point{3, 4}
	fmt.Println(p)
	p.MoveYTo(31)
	fmt.Println(p)
}

/*
พอ MoveYTo สิ่งที่มันรับเข้ามาคือ address of point
variable ตรงนี้มันก็คือ struct อันนึง แล้วมันก็ point มาที่ 3,4
ดังนั้นมันถึงเปลี่ยนค่าในนี้ได้ เพราะมันอ้างอิงตน.ที่ตั้งค่า X,Y ไม่ใช้ copy ค่า X,Y

ถ้า reciver ของเราเป็นชนิด pointer แล้วเราเอา pointer เนี่ยไปเรียก method ที่
รับ by value compiler จะทำการ insert (*) star ให้ ก็คือจะ derefernce ให้
พอ dereference ปุ๊ปมันก้จะวิ่งมาที่ solid object แล้ว solid object ก็จะ
copy ตัวเองไปให้ method ที่รับ by value ซึ่ง method ตรงนี้ก็ทำการเปลี่ยนแปลง
object ที่อยู่ตรงนี้ ไม่มีผลกระทบใดๆต่อต่อตัว นู้นเลย


*/
