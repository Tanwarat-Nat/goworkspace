package main

import (
	"fmt"
)

func main() {
	const (
		Sunday    = 1
		Monday    = 2
		Tuesday   = 3
		Wednesday = 4
		Thursday  = 5
		Friday    = 6
		Saturday  = 7
	)
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
	// 1 2 3 4 5 6 7
	const (
		Jan = 1
		Feb
		Mar
		Apr
		May
		Jun
	)
	fmt.Println(Jan, Feb, Mar, Apr, May, Jun)
	// 1 1 1 1 1 1

	const (
		Jul = iota // 0
		Aug
		Sep
		Oct
		Nov
		Dec
	)
	fmt.Println(Jul, Aug, Sep, Oct, Nov, Dec)
	// 0 1 2 3 4 5

	const (
		one float64 = iota + 1 //0+1 -> constant generator
		two
		three
		four
		five
	)
	fmt.Println(one, two, three, four, five)
	// 1 2 3 4 5
	// เป็น constant generator
	// ระบุ type = iota + number(int/float)ก็ได้

	const (
		six   float64 = iota + 1.1 // 0
		seven                      // 1
		eight                      // 2
		nine  int     = iota       // 3
		ten                        // 4
	)
	fmt.Println(six, seven, eight, nine, ten)
	// 1.1 2.1 3.1 3 4
	// เจอเงื่อนไขกัน ก็ผันตามคอนเซ็ปของเงื่อนไขนั้น

	const (
		ex float64 = iota + 1.1
		ex2
		ex3
		ex4 int = iota
		ex5
		ex6 = 99
		ex7
		ex8 = 99 + iota
		ex9
		ex10
	)
	fmt.Println(ex, ex2, ex3, ex4, ex5, ex6, ex7, ex8, ex9, ex10)
	// 1.1(0) 2.1(1) 3.1(2) 3 4 99 99 106(99+7) 107(99+8) 108(99+9)
}

// 1 -> ผลลัพธ์เป็น 1 ไล่ลงมาทุกตัว
// 1, 2, 3,..., n -> ผลลัพธ์เป็นตามที่กำหนด
// iota -> 0, 1, 2, 3, ...., n
// iota + (int หรือ float) ได้ -> 1,2,..,n | 1.2, 2.1,..n
// เจอเงื่อนไขกัน ก็ผันตามคอนเซ็ปของเงื่อนไขนั้น
