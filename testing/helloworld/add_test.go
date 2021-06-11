package main

import (
	"testing"
)

//func Test[Name] (*testing.T) {}

func TestAdd(t *testing.T) {
	a := add(1, 2)
	if a != 3 {
		t.Errorf("a := add(1,2) : is not 3. Got %d", a)
	}
}

func TestAdd2(t *testing.T) {
	a := add(2, 2)
	if a != 4 {
		t.Errorf("a := add(2,2) : is not 4. Got %d", a)
	}
}

/*
พอเรา ctrl clik add มันจะเด้งไปหา func add() มันจะหาว่า add() นี้มัน define
function ไว้ที่ไหนบ้าง ซึ่งอันนี้มัน define ไว้ที่ main.go ก็เป็นเหตุผลที่ว่าเราไม่จำเป็น
ต้อง redecla add() ไว้ในไฟล์ add_test.go จากนั้นเราจะเทสโดยใช้ comand go test

*/
