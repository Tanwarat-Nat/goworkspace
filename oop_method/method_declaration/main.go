package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
	// moveXTo float64 --> ห้าม! ประกาศ filed name ซ้ำกับ function(method)
}

type Vector struct {
	X, Y float64
}

func (p *Point) Length() float64 {
	return math.Hypot((*p).X, (*p).Y)

}

func (p *Point) MoveXTo(newX float64) {
	(*p).X = newX
}

func (p *Vector) moveXTo(newX float64) {
	(*p).X = newX
}

func (p *Point) moveYTo(newY float64) {
	(*p).Y = newY
}

func main() {
	p := Point{3, 4}
	fmt.Println("X, Y", p)
	fmt.Println(p.Length()) //5
	p.MoveXTo(6)
	fmt.Println("newX, newY :", p)
	fmt.Println("newLength :", p.Length())
}

/*
การสร้าง method ไม่ยากเลย การที่เราจะ associate method ตรงนี้
ไปไว้ที่ type ที่เป็น Point เนี่ยไม่ยาก

Method เป็นฟังก์ชั่น
Point is a Method's reciver
สมัยก่อน
ถ้าเราจะส่ง message เข้ามาหาฟังก์ชั่นที่ชื่อว่า moveXTo แล้วฟังก์ชั่นตรงนี้
p.MoveXTo(6) คนรับสารจากฟังก์ชั่นตรงนี้คือใคร
p = Point{3,4} ซึ่งในนี้ก็คือ p เราจึงเรียก Point ตรงนี้ p ตรงนี้ว่า Method's reciver
มันจะไม่มีคนอื่นมา recive ตรงนี้อีกแล้ว

การเชื่อมกันระหว่าง method name กับ base type/reciver type
ไม่มี this, self  -> this.X = newX ใช้แค่ Method's reciver name
เป็นการตอกย้ำคำศัพท์ที่ชื่อว่า reciver
p เป็น reciver จะมี name เป็นอย่างอื่นก็ได้ เช่น pg มี base type เป็น Point
p.MoveXTo(6) ตรงจุด (.) เรียกว่า selector เพราะ . มัน select method
เราจะประกาศฟังก์ชั่นที่มีชื่อซ้ำกับตัวแปรไม่ได้ เสือ 2 ตัวอยู่ถ้ำเดียวกันไม่ได้ ไม่ยอม
เพระามันไม่รู้ว่าจะไปเรียกใช้ตัวไหนกันแน่
แต่ถ้าประกาศตัว reciver มันต่างกัน เช่น Point กับ Vector
ชื่อ method มันจะซ้ำกันได้ไม่มีปัญหา เพราะว่ามันไม่ได้ผูกกับชื่ออย่างเดียว
มันผูก name กับ reciver type




*/
