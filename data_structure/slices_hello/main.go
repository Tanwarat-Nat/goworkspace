package main

import (
	"fmt"
)

func main() {
	fruits := [5]string{"apple", "banana", "papaya", "orange", "mango"}
	myFavor := fruits[1:4] // execute ที่ run time เพราะปรับย่อ-ขยายได้
	fmt.Println("my =", myFavor, len(myFavor), cap(myFavor))

	yourFavor := myFavor
	fmt.Println("your =", yourFavor, len(yourFavor), cap(yourFavor))

	fmt.Println("---> yourFavor[0] = 'guava'")
	yourFavor[0] = "guava"
	fmt.Println("your =", yourFavor, len(yourFavor), cap(yourFavor))
	fmt.Println("my  =", myFavor, len(myFavor), cap(myFavor))

}
