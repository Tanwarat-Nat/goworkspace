package hello

import "testing"

func TestHello(t *testing.T) {
	want := "Hello"
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want : %q", got, want)
	}
}

func TestHelloTh(t *testing.T) {
	want := "สวัสดี-"
	if got := HelloTh(); got != want {
		t.Errorf("HelloTh() = %q, want : %q", got, want)
	}
}

/*
เราไม่ได้อยู่ใน GOPATH จะไปเทสข้างนอกเรื่อยเปื่อยไม่ได้ ต้องเทสในที่ของมัน

Lenovo@DESKTOP-A96UHVC C:\Users\Lenovo\Desktop\helloapp
$ go test -v .

ok      _/C_/Users/Lenovo/Desktop/helloapp      0.130s
go command ไม่รู้ว่าจั imporrt ตรงนี้มาจากไหน ก็เลยสร้างเฟค import มาให้

helloapp ของเรายังไม่ใช่ go module ถ้าเราจะทำยังมี requement อีกอันคือ go mod
Lenovo@DESKTOP-A96UHVC C:\Users\Lenovo\Desktop\helloapp
$ go mod init example.com/helloapp

$ go test -v .
ok      example.com/helloapp/Desktop/helloapp   0.118s
ทุกอย่างรันผ่าน ไม่ต้องสร้างเฟคพาร์ทแล้ว
*/
