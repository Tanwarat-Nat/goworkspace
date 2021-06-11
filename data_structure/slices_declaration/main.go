package main

import "fmt"

func main() {
	/*
		Arrays
		[10]int{ }
		[...]int{ }

		Slices

		- เอาข้อมูลมาจาก arrays

		x := a[i:j]

		- เราไม่รู้ว่าจะเอาข้อมูลมาจาก arrays ตัวไหน
		- จึงประกาศมาก่อนแล้วใส่ข้อมูลทีหลัง

		make([]int, len)
		make([]int, len, cap)
		^
		make(ใส่ประเภท, ใส่จำนวน, ใสจำนวน)

	*/

	a := []int{1, 2, 3, 4} // 0,1,2,3
	s := a[1:3]
	ss := a[:3]
	sss := a[1:]
	ssss := a[:]
	fmt.Println("s  [1:3] =", len(s), cap(s), s)
	fmt.Println("ss [:3]  =", len(ss), cap(ss), ss)
	fmt.Println("sss [1:]  =", len(sss), cap(sss), sss)
	fmt.Println("ssss [:]  =", len(ssss), cap(ssss), ssss)

	// มีข้อมูลจุที่ 5 แต่พื้นที่เก็บได้ถึง 10
	m := make([]int, 5, 10)
	// ใน [] ถ้าใส่จน.เกิน len ที่ระบุ จะ error
	// ถ้าจะทำต้องใช้ append มาช่วย
	m[0] = 99
	fmt.Println("m =", len(m), cap(m), m)

	/*
		len  <  0    =  panic
		cap  <  len  =  panic
		memory allocation -> fail  = stop
		ใช้ len เกินขีดจำกัดของ พื้นที่ memory
	*/

	// j := make([]int, -1, 0)
	// k := make([]int, 11, 10)
	i := make([]int, 999999999999)
	fmt.Println(i)
}
