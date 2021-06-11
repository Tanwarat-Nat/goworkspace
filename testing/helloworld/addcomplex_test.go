package main

import (
	"testing"
)

//func Test[Name] (*testing.T) {}

func TestAddComplex(t *testing.T) {
	a := add(1, 2)
	if a != 3 {
		t.Errorf("a := add(1,2) : is not 3. Got %d", a)
	}
}

func TestAddComplex2(t *testing.T) {
	a := add(2, 2)
	if a != 4 {
		t.Errorf("a := add(2,2) : is not 4. Got %d", a)
	}
}

/*
มีไฟล์เทสหลายอัน อยากเทสแค่ไฟล์นี้ ให้ใส่ regular expression
$ go test -v -run="TestAddComplex" testing/helloworld
regular expression มันจะไปจับกับชื่อฟังก์ชั่นมันไม่ได้มาจับกับไฟล์
ชื่อไม่ต้องเป๊ะ 100% แค่มีส่วนนึงในชื่อก็ได้แล้ว เช่น "TestAdd" ก็จะจับหมดเลย
4 function จะเป้นชื่อตัวท้ายก็ได้ เช่น "lex"

regular expression มันจะมา match กับ function name

*/
