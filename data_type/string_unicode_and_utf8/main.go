package main

import "fmt"

func main() {
	x := "hi-สวัสดี" // []byte{0,1,...,21}
	y := []byte{0x68, 0x69, 0x2d, 0xe0, 0xb8, 0xaa, 0xe0, 0xb8, 0xa7, 0xe0, 0xb8, 0xb1, 0xe0, 0xb8, 0xaa, 0xe0, 0xb8, 0x94, 0xe0, 0xb8, 0xb5}
	z := []rune(x)
	fmt.Println(x, len(x))
	fmt.Println("len(y) =", len(y))
	fmt.Println("len(z) =", len(z))
	fmt.Printf("%q\n", z[5])

	fmt.Println(string(y[0]))                 // h
	fmt.Println(string(y[1]))                 // i
	fmt.Println(string(y[2]))                 // -
	fmt.Println(string(y[3]), string(y[3:6])) // ส
	fmt.Println(string(y[4]), string(y[6:9])) // ว
	fmt.Println(string(0xE2A))                // unicode In hex U+0E2A
	fmt.Println("utf-8", "\xe0\xb8\xaa")      // ส

	// 1110 0010 1010             (ส) unicode -> utf8
	// 1110xxxx	10xxxxxx  10xxxxxx
	// 11100000 10111000  10101010 = e0b8aa
}

// https://charnnarong.github.io/digitconverter/
// https://unicodelookup.com/
// https://en.wikipedia.org/wiki/UTF-8#:~:text=UTF%2D8%20is%20a%20variable,Transformation%20Format%20%E2%80%93%208%2Dbit.

// ส - Unicode In hex      U+0E2A
// ส - Unicode In binary   1110 0010 1010

// unicode ใช้ space มากเกินไป
// 1 unicode ใช้พื้นที่ 4 byte =
// (4 byte) * (8 bit) = 32 bit
//
// เปลี่ยน unicode -> utf8
