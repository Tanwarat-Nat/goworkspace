package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	x := complex128(complex(3, 2))
	var y complex128 = complex(1, 7)

	fmt.Println("x + y", x+y)
	fmt.Println("x * y", x*y)
	fmt.Println(cmplx.Sqrt(-9))
	fmt.Println(cmplx.Abs(complex(3, 4)))
	fmt.Println(cmplx.Abs(complex(3, -4)))
}

// (3 + 2i) + (1 + 7i)
//  3 + 1 + (2 + 7)i
//  4 + 9i
