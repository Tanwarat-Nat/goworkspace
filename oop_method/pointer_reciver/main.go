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
	(&p).MoveYToByValue(31) // (*(&p)) compile insert * (derefernce) -> p
	fmt.Println(p)
}

/*
recap
pointer reciver เรียกใช้ method value reciver
compiler จะทำการ insert (*) คือจะ derefernce
solid object จะทำการ copy ตัวเองไปให้ method
สุดท้ายก็ไม่มีอะไรเปลี่ยนเหมือนเดิม

sub
ถ้า reciver ของเราเป็นชนิด pointer แล้วเราเอา pointer เนี่ยไปเรียก method ที่
รับ by value compiler จะทำการ insert (*) star ให้ ก็คือจะ derefernce ให้
พอ dereference ปุ๊ปมันก้จะวิ่งมาที่ solid object แล้ว solid object ก็จะ
copy ตัวเองไปให้ method ที่รับ by value ซึ่ง method ตรงนี้ก็ทำการเปลี่ยนแปลง
object ที่อยู่ตรงนี้ ไม่มีผลกระทบใดๆต่อต่อตัว นู้นเลย


*/
