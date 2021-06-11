package main

import "fmt"

func main() {

	// switch expression {
	// case condition:

	// }

	x := 1
	switch x {
	case 1:
		fmt.Println("Xone")
	case 2:
		fmt.Println("Xtwo")
	case 3:
		fmt.Println("Xthree")
	default:
		fmt.Println("X in default")
	}

	y := 1
	switch y {
	case 1:
		fmt.Println("Yone")
		fallthrough // จะไปจบการทำงานที่ 2 two
	case 2:
		fmt.Println("Ytwo")
		fallthrough // จะไปจบการทำงานที่ 3 three
	case 3:
		fmt.Println("Ythree")
	default:
		fmt.Println("Y in default")
	}

	z := 200
	switch true { // ถ้า omit expression ก็เป็น true เหมือนกัน
	case 1 <= z && z < 100:
		fmt.Println("Zdeci")
	case 100 <= z && z < 1000:
		fmt.Println("Zcenti")
	case 1000 <= z:
		fmt.Println("Zmeter")
	default:
		fmt.Println("Z in default")
	}

	i := 200
	switch false {
	case 1 <= i && i < 100:
		fmt.Println("Ideci")
	case 100 <= i && i < 1000:
		fmt.Println("Icenti")
	case 1000 <= i:
		fmt.Println("Imeter")
	default:
		fmt.Println("I am default")
	}

	zi := -20
	switch true {
	case 1 <= zi && zi < 100:
		fmt.Println("Ideci")
	case 100 <= zi && zi < 1000:
		fmt.Println("ZIcenti")
	case 1000 <= i:
		fmt.Println("ZImeter")
	default:
		fmt.Println("ZI am default")
	}

	j := "b"
	switch j {
	case "a":
		fmt.Println("jA")
	case "b":
		fmt.Println("jB")
	case "c":
		fmt.Println("jC")
	case "d", "e", "f", "g":
		fmt.Println("jD", "jE", "jF", "jG")
	default:
		fmt.Println("j in default")
	}

	k := "g"
	switch k {
	case "a":
		fmt.Println("kA")
	case "b":
		fmt.Println("kB")
	case "c":
		fmt.Println("kC")
	case "d", "e", "f", "g":
		fmt.Println("kD", "kE", "kF", "kG")
		if k == "g" {
			break
		} // ถ้า k := "g" จะ break และไม่ print lucky ออกมา
		fmt.Println("k in lucky")
	default:
		fmt.Println("k in default")
	}

	a := "z"
	switch a {
	case "a":
		fmt.Println("aA")
	case "b":
		fmt.Println("aB")
	case "c":
		fmt.Println("aC")
	case "d", "e", "f", "g":
		fmt.Println("aD", "aE", "aF", "aG")
		if k == "g" {
			break
		}
		fmt.Println("a in lucky")
	case "z": // z ไม่ทำอะไรสักอย่าง คือไม่ต้องทำไรเลย ไม่ต้องprint
	default:
		fmt.Println("a in default")
		//case "z": // วางตรงไหนก็ได้
	}

	// var b interface
	// b := "f"
	// switch b {
	// case int:
	// 	fmt.Println("A")
	// case float64:
	// 	fmt.Println("B")
	// case string:
	// 	fmt.Println("C")
	// case "d", "e", "f", "g":
	// 	fmt.Println("D", "E", "F", "G")
	// 	if k == "g" {
	// 		break
	// 	}
	// 	fmt.Println("You are lucky")
	// default:
	// 	fmt.Println("I am default")
	// }
}

// condition ถูก apply ทำ case นี้ แล้วจบเลยไม่มี break
// ภาษา C จะ fallthrough ไปเรื่อยๆจนกว่าจะเจอ break
// ไม่จำเป็นต้องใช้ break ในภาษ Golang
