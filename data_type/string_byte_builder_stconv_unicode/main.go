package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	// byte package
	x := "ทดสอบ"
	finder := "สอ"
	fmt.Println()
	fmt.Println(bytes.Count([]byte(x), []byte(finder)))
	// ไ้ด้เหมือนกัน
	fmt.Println(strings.Count(x, finder))

	//buff := bytes.Buffer{}
	buff := strings.Builder{}
	buff.WriteRune('h')
	buff.WriteRune('i')
	fmt.Println(buff)
	fmt.Println(buff.String())

	// stconv
	atoi, _ := strconv.Atoi("123")
	fmt.Println(atoi, reflect.TypeOf(atoi))

	itoa := strconv.Itoa(123)
	fmt.Println(itoa, reflect.TypeOf(itoa))

	fmt.Println(strconv.ParseBool("True"))
	fmt.Println(strconv.ParseBool("true"))
	fmt.Println(strconv.ParseBool("1"))
	fmt.Println(strconv.ParseBool("False"))
	fmt.Println(strconv.ParseBool("fALse"))
	fmt.Println(strconv.ParseBool("0"))

	fmt.Println(strconv.ParseBool("trUE"))
	fmt.Println(strconv.ParseBool("5"))

	// unicode
	fmt.Println("unicode.IsDigit('1') : ", unicode.IsDigit('1'))
	fmt.Println("unicode.IsDigit('๒') : ", unicode.IsDigit('๒'))
	fmt.Println("unicode.IsDigit('四') : ", unicode.IsDigit('四'))
	fmt.Println("unicode.IsNumber('四') : ", unicode.IsNumber('四'))
	fmt.Println("unicode.IsUpper('a') : ", unicode.IsUpper('a'))
	fmt.Println("unicode.IsUpper('a') : ", unicode.IsUpper('a'))
	fmt.Println("unicode.IsUpper('A') : ", unicode.IsUpper('A'))
	fmt.Println("unicode.In('ฟ', unicode.Thai) : ", unicode.In('ฟ', unicode.Thai))
}
