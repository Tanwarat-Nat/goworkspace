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

pointer reciver เรียกใช้ method pointer reciver
ก็ไม่จำเป็นจะต้อง class อะไร สื่งที่มัน ก๊อปไปก็คือ structure (pointer) ตรงนี้
ไปไว้ให้ method pointer ซึ่ง structure ตัวนี้ มันก็ยัง point
มาที่ solid object ตน.เดิมอยู่ดี
ดังนั้น method pointer reciver มันก็เลยทำการเปลี่ยนแปลง solid object ตรงนี้ได้
เพราะว่า point ไปที่ตน.เดียวกัน


ถ้าพาส p เฉยๆ ตรงนี้มันเป็น solid object
compiler มันทำการ insert (&) address ให้เราอัตโนมัต

*/
