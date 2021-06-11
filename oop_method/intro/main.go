package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func Length(p *Point) float64 {
	return math.Hypot((*p).X, (*p).Y)

}

func MoveXTo(p *Point, newX float64) {
	(*p).X = newX
}

func moveYTo(p *Point, newY float64) {
	(*p).Y = newY
}

func main() {
	p := Point{3, 4}
	fmt.Println("X, Y", p)
	fmt.Println("Length :", Length(&p))
	MoveXTo(&p, 6)
	fmt.Println("newX, newY :", p)
	fmt.Println("newLength :", Length(&p))
}

/*
เขียนโค๊ดให้อยู่รูปเชิง oop  ในภาษาโก เทคนิคแบบไหนที่สามารถตอบโจทย์ oop feature ได้
ในโก ก็จะมีลิมิตต่างจากภาษอื่นๆ การจะสร้าง object อันนึง ในภาษาโกเราต้องใช้ struct
เพราะภาษาโก ไม่มีคำว่า class คีเวิร์ดอย่างในภาษาอื่น อย่าง จาวา จาวาสคริป
ถ้าใาจากภาษาซี struct จะเป็นคีเวิร์ดแรกๆที่เราเจอก่อนที่เราจะไปเจอคำว่า class ใน c++

struct ในภาษา GO มีความสามารถมากกว่า struct ในภาษา C แล้วสามารถตอบโจทย์
การเขีบนโปรแกรมในเชิง oop ได้ด้วย

การประกาศฟังก์ชั่นแบบไม่มี body คือการเตรียม signature ไว้ เพื่อ body ที่ถูกสร้างในภาษาอื่น

method คือ ฟังก์ชั่นที่ refer ถึง type ตรงนี้ได้
*/
