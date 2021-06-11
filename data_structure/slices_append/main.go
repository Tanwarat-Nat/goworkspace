package main

import "fmt"

func main() {
	x := [10]int{}
	a := x[0:5] // [:5]
	b := x[5:7]
	for i := range a {
		a[i] = i * i
	}
	b[0] = 98
	b[1] = 99
	// slice = append(slice, anotherSlice...) เอามาจาก slice อื่น
	a = append(a, b...) // a = append(a, b[0], b[1])
	a[len(a)-1] = 25    // b จะเปลี่ยนด้วย -> 98, 25
	// slice = append(slice, elem1, elem2) ต้อง type เดียวกันด้วย
	a = append(a, 71, 72, 73) // ยังเติมได้อยู่เพราะพท.เหลือ
	a = append(a, 74)
	/*เมื่อพื้นที่ไม่พอ x ก็ยังคงเดิมแหละ แต่ allocate พื้นที่ใหม่เป็น
	unnamed จะ refer มันได้ผ่าน slice a จากนั้นจะ copy ข้อมูลทั้งหมดมาไว้
	แล้ว append 71,72,73,74 เข้ามา แล้วไปเพิ่ม capacity ที่ a
	จะเพิ่มให้เท่าไหร่ขึ้นอยู่กับ algorithm
	*/
	fmt.Println("x :\t", x) // x จะ update ตามตลอด
	fmt.Println("a :\t", len(a), cap(a), a)
	fmt.Println("b :\t", len(b), cap(b), b)

}
