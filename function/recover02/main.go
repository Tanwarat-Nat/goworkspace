package main

import "fmt"

func A(x int) {
	defer func() {
		fmt.Println("defer in A")
		if p := recover(); p != nil {
			panic(fmt.Sprintf("A: %v", p))
		}
	}()
	B(x)
	fmt.Println("hello in A")
}

func B(x int) {
	defer func() {
		fmt.Println("defer in B")
		if p := recover(); p != nil {
			panic(fmt.Sprintf("B: %v", p))
		}
	}()
	C(x)
	fmt.Println("hello in B")
}

func C(x int) {
	defer func() {
		fmt.Println("defer in C")
		if p := recover(); p != nil {
			panic(fmt.Sprintf("c: %v", p))
		}
	}()
	if x == 4 {
		panic("for no reason")
	}
}

func main() {
	A(4)
}

/*
usecase ที่เราจะทำกันจริงๆก็คือ ถ้าเกิดว่ามันมี panic เกิดขึ้นเราจะพยายามไม่ handle มัน
เราจะพยายาม retrow panic ออกไป เพราะรู้ว่า panic มันเกิดจาก C
แต่ว่า แบบนี้เราจะเจอ for no reason เราไม่รู้มันเกิดจากที่ไหน เราพยายามจะ enlist message
ตรงนี้ ที่ B ก็เหมือนกัน ถ้ามันเกืิด panic ก็ รีเทอร์นตรงนี้เป็นบี

น้อยมากที่จะใช้ panic  พอใช้ปุ๊บเราก็พยายามให้ panic มีข้อความชัดเจนมากขึ้น

*/
