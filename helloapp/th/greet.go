package th

import "github.com/Tanwarat-Nat/quote/v2"

func Greet() string {
	return "สวัสดี-"
}

func Greet2() string {
	return "สวัสดี-" + quote.Say2()
}

func Greet3() string {
	return "สวัสดี-" + quote.Speak2()
}

/*
 เราจะเปลี่ยน package eng ไปเป็น module เพื่อว่าอีกสักพักเราต้องใช้ dependency
 เวอร์ชั่นต่างกันกับอันอื่น

 มันก็จะกลายเป็นมี module ข้างนอก มี module ซ้อนเข้าไปอีก ก็คือ eng
 Lenovo@DESKTOP-A96UHVC C:\Users\Lenovo\Desktop\helloapp\eng
$ go mod init exmple.com/helloapp/eng

*/
