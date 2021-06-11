package main

import "fmt"

func createF(list []int) []func() {
	result := []func(){}

	for _, v := range list {
		captureV := v
		fmt.Println("inside createF", captureV)
		result = append(result, func() {
			fmt.Println(captureV)
		})
	}
	return result
}

func main() {
	fList := []func(){} // [ f1(), f2(), f3() ]
	fList = createF([]int{108, 12, 30, 40, 10})
	for _, v := range fList {
		v()
	}
}

/*
anonymous ภายใต้ loop
*/
