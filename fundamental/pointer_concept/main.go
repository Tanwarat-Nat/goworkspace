package main

import "fmt"

func main() {

	x := 70
	var p *int = &x

	fmt.Println("value(x): ", x)
	fmt.Println("address(&x): ", &x)
	fmt.Println("value/pointer(p = &x) : ", p)
	fmt.Println("address(&p) : ", &p)
	fmt.Println("value(*p) : ", *p) //dereference

}

// int float string เป็น value type
// ทำให้เป็น reference type โดยใส่ *type
