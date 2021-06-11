package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {

	var w io.Writer //zero value -> nil
	fmt.Printf("w = io.Writer: %T, %v\n", w, w)

	w = os.Stdout
	fmt.Printf("w = os.Stdout: %T, %v\n", w, w)
	fmt.Println("os.Stdout == nil:", w == nil)

	w = &bytes.Buffer{}
	w.Write([]byte("hello"))                          // ถ้าเราจะเขียนก็ทำแบบนี้ได้
	fmt.Printf("w = &bytes.Buffer{}: %T, %v\n", w, w) //value ของมันไม่มีอะไร

	w = nil
	fmt.Printf("w = nil: %T, %v\n", w, w)
	var byteBuff io.Writer //-> ไม่ก็ประกาศแบบนี้ไปเลยก็จะไม่แพนิค
	//var byteBuff *bytes.Buffer
	//เพื่อไม่อยากให้แพนิค ค่าตรงนี้ก็ไม่ควรเป็น nil สร้างอะไรให้มัน concrete type ขึ้นมาอันนึง
	byteBuff = &bytes.Buffer{}
	printByff(byteBuff)
	// fmt.Println("byteBuff == nil: ", byteBuff == nil)
	// fmt.Println("w == nil:", w == nil)

	// w = byteBuff
	// fmt.Println("w = byteBuff:", w == nil)
	// fmt.Printf("w = byteBuff: %T, %v\n", w, w)
	// พอ w เป็น nil แล้ว เราก็ให้ w = byteBuff ซึ่ง byteBuff คือ *bytes.Buffer
	// ตัวนี้มันก็ทำการ implement io.Writer เหมือนกัน ดังนั้น type ของ w ที่เกิดขึ้น
	// จะเป็นแบบนี้ *byte.Buffer แต่ว่า value มันเป็น nil นั่นก็เป็นเหตึุผลว่าทำไม
	// byteBuff == nil
	/*ถ้าเราจะ check ความเป็น nil ของ interface ถ้าจะทำแบบ w == nil ให้ระลึกว่าการทำ
	แบบนี้ จะเป็น nil ได้ก็ต่อเมื่อ type กับ value มันเป็น nil เท่านั้น*/

}

func printByff(w io.Writer) {
	if w != nil {
		fmt.Println(w.Write([]byte("hello")))
	}
}

// printByff(byteBuff) พอ run จะเกิด panic เลย มันเข้ามาได้ไง ตรงนี้มันเป็น nil นะ
// ที่เข้ามาได้เพราะ มันไม่ได้เช็คแค่เฉพาะ value เท่านั้น มันเช็ค type ด้วย

/*
interface มี value มี type ของมันอยู่ การไม่เข้าใจ type กับ value ของ interface
ทำให้เราอ่าน code ของคนอื่นไม่เข้าใจ ทำให้บางทีเราเขียนโค๊ดของเราเองแล้วก็เกิดปัญหา
แล้วก็ไม่รู้ว่าผิดตรงไหน

ใน interface ที่ประกาศมา มันจะมีข้อมูลอยู่ 2 อัน

io.Writer -> interface
Type	nil
Value	nil

zero value ของ interface type กับ value จะเป็น nil

ต่อไป ถ้าเราเอา interface ตรงนี้ไปรับค่า ที่เป็น concret type ที่ทำการ implement
io.Writer ตยใก็คือ os.Stdout สิ่งที่เกิดขึ้นคือ type มันจะเปลี่ยนไป

io.Writer
Type	*os.File
Value	&(0xc000082280)

มันไม่ได้มีแค่ os.Stdout มันยังมี interface ที่เป็น standard package อันนึงที่เราใช้กันบ่อยๆ
นั่นก็คือ &bytes.Buffer{}  Buffer ภายใต้ bytes package ก็ทำการ implement io.Writer
เหมือนกัน ถ้าเราลองทำแบบนี้ดู type ของเราก็จะเปลี่ยนจาก *os.File -> *bytes.Buffer

io.Writer
Type	*bytes.Buffer
Value	<zero value buffer>

และถ้าเราให้ io.Writer เป็น nil ทั้งหมดก็จะเปลี่ยนเป๋น nil

io.Writer
Type	nil
Value	nil

ถ้าเราจะ check ความเป็น nil ของ interface ถ้าจะทำแบบ w == nil ให้ระลึกว่าการทำ
แบบนี้ จะเป็น nil ได้ก็ต่อเมื่อ type กับ value มันเป็น nil เท่านั้น


*/
