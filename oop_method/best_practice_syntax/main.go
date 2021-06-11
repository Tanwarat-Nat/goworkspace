package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func (p *Point) Length() float64 {
	return math.Hypot(p.X, p.Y) // compiler insert * -> (*p).X, (*p).Y
}

func (p *Point) MoveXTo(newX float64) { // หน้าตาแบบนี้ คือ method
	p.X = newX // compiler auto insert * -> (*p).X
}

func (p *Point) MoveYTo(newY float64) {
	MoveYTo(p, newY)
}

func MoveYTo(p *Point, newY float64) { // หน้าตาแบบนี้ คือ function
	p.Y = newY // compiler auto insert * -> (*p).Y
}

func main() {
	p := Point{3, 4}
	fmt.Println(p)
	p.MoveYTo(31)
	fmt.Println("p.MoveYTo(31):", p)

	//bytes.Buffer
}

/*
ถ้าเราสร้าง method ให้ struct ใดๆแล้วเนี่ย ถ้า method ใด method นึง
มี reciver ชนิดเป็น pointer ของ struct นั้น method ที่เหลือก็จะต้องรับเป็น
pointer ของ struct ตรงนี้เหมือนกัน ถึงแม้ว่า method ตัวนั้นจะไม่ต้องการเปลี่ยนแปลง
ข้อมูลที่อยู่ข้างในก็ตาม

แต่เรื่อง compiler ช่วยเราในเรื่องของ syntax ใช้ไม่ได้กับ function argrement

*/
