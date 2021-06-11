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
	p.MoveYTo(31) // ถ้าใส่ p เฉยๆ compiler จะ insert (&) address ให้
	fmt.Println(p)
}

/*
solid object -> p
เรียก method value reciver  -> MoveYToByValue
ก็ไม่ต้องทำอะไร ไม่มีอะไรเปลี่ยนอยู่แล้ว

แต่ถ้า
solid object -> p
เรียก method pointer reciver
compiler ก็จะ insert (&) ให้อัตโนมัต
*/
