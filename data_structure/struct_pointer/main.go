package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	x := Person{"Filicity", 23}
	y := &x
	y.name = "Oliver" // (*y).name = "Oliver"

	fmt.Println(x)
	fmt.Println(*y) // ใส่ * เพื่อเอา & ออกจากตัวผลลัพธ์

	i := Person{"Filicity", 23}
	z := Person{"Filicity", 23}
	j := &Person{"Filicity", 23}

	fmt.Println(i == z)    //ต้องเท่ากันทุกกระเบียดนิ้วถึง true
	fmt.Printf("%v\n", i)  // defualt format
	fmt.Printf("%+v\n", z) // %+v จะใส่ flied name ให้ด้วย
	fmt.Printf("%+v, %T\n", *j, j)

	n := new(Person)
	*n = Person{"Filicity", 23}
	fmt.Println(*n)

	m := new(Person)
	m.name = "Filicity" // (*m).name = "Filicity"
	m.age = 23
	fmt.Println(*m)

}

// ถ้าระบุ filed name ไว้ด้วย จะบอกไม่ครบก็ได้ ผลจะแสดงค่า zero value แทน
// แต่ถ้าบอกครบ จะระบุ filed name ไม่ครบ ไม่ได้ ต่องระบุให้ครบ

// struct copy by value ถ้าเปลี่ยน y แล้ว x ก็จะไม่เปลี่ยน
// ถ้าจะให้เปลี่ยนก็ใส่ &x ให้เป็นค่า pointer
