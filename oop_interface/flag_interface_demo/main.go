package main

import (
	"flag"
	"fmt"
	"strconv"
)

// type Value interface {
// 	String() string
// 	Set(string) error
// }

var romanMap = map[string]int{
	"i":   1,
	"ii":  2,
	"iii": 3,
	"iv":  4,
	"v":   5,
}

type RomanAge struct {
	age string // i = 1, ii = 2, v = 5
}

func (r *RomanAge) String() string {
	return strconv.Itoa(romanMap[r.age])
}

func (r *RomanAge) Set(value string) error {
	r.age = value
	return nil
}

func flagRomanAge() *RomanAge {
	romanAge := RomanAge{}
	flag.Var(&romanAge, "age", "help message for flagname")
	return &romanAge
}

var name = flag.String("name", "anynymous", "your name")

var age = flagRomanAge()

func main() {
	flag.Parse()
	fmt.Println(*name)
	fmt.Println(age)

}

// $ go run main.go -name Filicity -age iii
// Filicity
// 3

/*
การใช้ interface อีกแบบที่ popular มากในการเขียนโปรแกรมที่รับ flag เข้ามา
โปรแกรมที่รับ flag เข้ามายกตัวอย่างเช่น ถ้าสมมมติว่าเรารัน flag
เราไม่ต้องการ fic ว่าต้องเป็นชื่อ เดวิด เราต้องการให้เป็นชื่อคนอื่น
/flaf -name Filicity --> แบบนี้เรียกว่า flag agrement ว่ามีอะไรบ้าง
นอกจาก name อาจมีอายุ หรือที่อยู่

fmt.Println(os.Args[1:])

$ go run main.go -name Filicity -age 1234
[-name Filicity -age 1234]

แบบนี้มันไม่กรุ๊ปกัน มันไม่ได้เอา 1234 กรุ๊ปกับอายุ ไม่ได้เอา Filicity ไปกรุ๊ปกับ name
เราต้องการกรุ๊ป เราไม่ได้อยากมา split ใน slice ของ string แล้วก็มา match กันเอง
เราก็เลยมาเรียกใช้ package flag.Parse()

*/
