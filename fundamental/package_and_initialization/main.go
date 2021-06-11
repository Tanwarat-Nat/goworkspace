package main

import (
	"fmt"
	"math/rand"
	"prime"
)

func main() {
	for i := 0; i <= 100; i++ {
		x := rand.Intn(1000000)
		fmt.Printf("%d , %t\n", x, prime.IsPrime(x))
	}

	fmt.Println(prime.IsPrime(2))
	fmt.Println(prime.IsPrime(3))
	fmt.Println(prime.IsPrime(4))
	fmt.Println(prime.IsPrime(390342))
	fmt.Println(prime.IsPrime(13))
}
