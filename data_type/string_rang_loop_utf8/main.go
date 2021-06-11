package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// rang loop
	x := "ทดสอบ"
	for i, v := range x {
		//fmt.Println(i, v, reflect.TypeOf(v))
		fmt.Printf("%d, %c\n", i, v)
	}

	// utf-8 package
	fmt.Println("utf8.RuneCountInString(x): ", utf8.RuneCountInString(x))

	for i := 0; i < len(x); {
		r, s := utf8.DecodeRuneInString(x[i:])
		i += s
		fmt.Printf("%c ", r)
	}
}
