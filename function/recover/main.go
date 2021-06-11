package main

import "fmt"

func A(x int) {
	defer func() {
		fmt.Println("defer in A")
	}()
	B(x)
	fmt.Println("hello in A")
}

func B(x int) {
	defer func() {
		fmt.Println("defer in B")
	}()
	C(x)
	fmt.Println("hello in B")
}

func C(x int) {
	defer func() {
		fmt.Println("defer in C")
		if p := recover(); p != nil {
			fmt.Println("handling panic:", p)
		} //
	}()
	if x == 4 {
		panic("for no reason")
	}
}

func main() {
	A(4)
}

/*
ต่อจาก panic04

ถ้าเราอยากให้โปรแกรมจบ panicking state พอ panicking state เริ่มต้นที่ C แล้ว
เราอยากจะให้มันจบที่ B ถ้ามันจบที่ B  เสร็จดังนั้น B ตัวนี้ก็จะไม่ panicking state
พอไม่เป็น panicking state มันก็จะสามารถปริ้นได้

ต่อจากคลิปที่แล้ว โปรแกรมของเราจะเกิด panic ที่ function C
แล้ว panic ก็จะ bubble มาหา function B bubble มาหา  function A
เราเลยไม่เห็น fmt.Println("hello in A")

สมมติว่าถ้า function C ไม่เกิด panic หมายความว่าโปรแกรมจะ execute ในบรรทัดถัดไป
นั่นก็คือ B

การ handle panic ที่เกิดขึ้น สามารถมาทำได้ เช่น ในขณะที่เราบอกว่า
	defer func() {
		fmt.Println("defer in C")
		if p := recover(); p != nil {
			fmt.Println("handlingh panic:",p)
		}
	}()
ตรงนี้เราสามารถเปลี่ยน state จาก panicking state ไปเป็น state ธรรมดาได้
นั่นก็คือเราใช้คำสั่ง recover
recover จะ return value ที่เรา pass ใน panic มา
หมายความว่า C ของเราจะไม่เกิด panic ดังนั้นเราควรจะเห็น "hello in B"
"hello in A" เพระาว่า B ก็ไม่เกิด panic เช่นเดียวกัน

recover ต้องประกาศใน defer function call
ถ้าเรียกข้างนอกมันขจะไม่เิกดอะไรขึ้น มันก็จะ return nil
เพระาว่ามันไม่มี panic ที่จะ return กลับมา
มันจะมีผลก็เมื่อเราเรียกที่ defer เพระาว่าถ้ามันเกิด panic defer มันก็จะต้องทำงาน


*/
