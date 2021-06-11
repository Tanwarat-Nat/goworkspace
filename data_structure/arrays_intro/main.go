package main

import (
	"fmt"
	"reflect"
)

func main() {
	/*

		การประกาศ Arrays

		[10]int{ }
		^
		var a [n]int{}

		[...]int{ }
		^
		name := [...]string{} -> ellisis

	*/

	// var a [n]int{}
	fruits := [5]string{
		"apple",
		"banana",
		"papaya",
		"orange",
		"mango",
	}
	fmt.Println(len(fruits), fruits)

	twoSlots := [2]int{}
	threeSlots := [3]int{}
	fmt.Println("Two slots\t", reflect.TypeOf(twoSlots))
	fmt.Println("Three slots\t", reflect.TypeOf(threeSlots))
	twoSlots = [2]int{1, 2}
	//	fmt.Println(twoSlots == threeSlots) // mismatch type
	// size ต่างกันการ assign ให้กันมันจะไม่ work

	// name := [...]string{} -> ellisis
	animal := [...]string{
		"dog",
		"cat",
		"rabbit",
		"lion",
		"elephant",
	}
	fmt.Println("animal =", animal)

	// indexName := [...]string{}
	// {INDEXx : value, INDEXy : value}
	number := [...]string{
		4:       "one", // index 4
		"two",   // index 5
		"three", // index 6
		"four",  // index 7
		"five",  // index 8
	}
	fmt.Println(len(number), number, number[4], number[5])

	number2 := [...]string{
		3:  "one",
		1:  "two",
		0:  "three",
		40: "four",
		2:  "five",
	}
	fmt.Println(len(number2), number2) //len = max+1

	/*  การเท่ากันของ array
	สลับตำแหน่งแค่อันเดียวก็จะเป็น false
	*/
	fmt.Println([2]int{1, 2} == [2]int{1, 2})
	fmt.Println([3]int{1, 2, 3} == [3]int{3, 1, 2})

	/*	 Coppy by value!!!!
	ถ้าเรา assign มันจะ coppy แต่ละ element ไปเลย
	เรามี array ตั้งต้นอันนึก แล้วมา assign มาให้อีกอันนึง
	ถ้าเราเปลี่ยนอันนี้ อันนั้นก็จะไม่เปลี่ยน เพราะมันคนละ space กันแล้ว
	*/
	fruit := [5]string{
		"apple",
		"banana",
		"papaya",
		"orange",
		"mango",
	}
	dubFruits := fruit // มันจะไม่แชร์ memory มันจะ coppy มาวางเลย
	fmt.Println(fruit)
	fmt.Println(dubFruits)

	fmt.Println("dubFruits[0] = 'watermelon'")
	dubFruits[0] = "watermelon"
	fmt.Println(fruit)
	fmt.Println(dubFruits)

	fmt.Println("ptrFruits := &fruit")
	ptrFruits := &fruit
	fmt.Println(fruit)
	fmt.Println(*ptrFruits)

	fmt.Println("ptrFruits[0] = 'watermelon'")
	ptrFruits[0] = "watermalon"
	fmt.Println(fruit)
	fmt.Println(*ptrFruits)

	// name := [มีกี่ปีกกา][แต่ละปีกกามีกี่ตัว]type

	a := [2]int{1, 2}
	b := [2][2]int{{1, 2}, {3, 4}}
	c := [2][3]int{{1, 2, 3}, {4, 5, 6}}

	fmt.Println("a =", a)
	fmt.Println("b =", b)
	fmt.Println("c =", c)

}
