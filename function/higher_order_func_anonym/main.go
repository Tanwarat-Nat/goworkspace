package main

import "fmt"

func createFib() func(int) []int {
	fList := []int{0, 1, 1, 2, 3, 5}
	return func(n int) []int {
		if n > len(fList) {
			for n > len(fList) {
				lastIndex := len(fList) - 1
				fList = append(fList, fList[lastIndex]+fList[lastIndex-1])
			}
		}
		return fList[:n]
	}
}

func main() {
	fib := createFib()
	fmt.Println(fib(5))
	fmt.Println(fib(6))
	fmt.Println(fib(60))
}

/* เป็น order ที่รับ function เข้ามาแล้วทำอะไรสักอย่างนึง
แล้วก็ return function ออกไป ซึ่งอาจจะเป็น function
ที่มีความสามารถมากกว่าที่ส่งเข้ามาก็ได้
*/
