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

func MoveYTo(p *Point, newY float64) { // หน้าตาแบบนี้ คือ function
	(*p).Y = newY
}

func main() {
	p := Point{3, 4}
	fmt.Println(p)
	p.MoveYToByValue(31)
	fmt.Println(p)
}

/*
พอมันทำมาถึง p := Point{3, 4} มันก็จะ copy ค่า X: 3, Y: 4 ไว้
พอมาถึง p.MoveYToByValue(31) มันก็ไปเรียก MoveYTo(&p, newY) ต่อ
พอมันจะเขากระบวนการ มันก็เห้ยย มีค่ามาให้อยู่แล้วหนิ จากที่ copy มาอ่ะ
จะให้เปลี่ยนอะไรละ ก็แสดงค่านี้กลับไปนั่นแหละ ค่าใหม่ที่แทนเข้ามา ก็ไม่ใช้ละ
ไม่มีที่จะให้ลงหนิ สุดท้ายก็ไม่มีอะไรเปลี่ยน*/
