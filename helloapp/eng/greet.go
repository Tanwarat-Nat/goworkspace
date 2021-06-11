package eng

import "github.com/Tanwarat-Nat/quote"

func Greet() string {
	return "Hello"
	//	return quote.Speak() // "Hello"
}

func Greet2() string {
	return "Hello, " + quote.Say()
}

func Greet3() string {
	return "Hello, " + quote.Talk()
}

/*
 เราจะเปลี่ยน package eng ไปเป็น module เพื่อว่าอีกสักพักเราต้องใช้ dependency
 เวอร์ชั่นต่างกันกับอันอื่น

 มันก็จะกลายเป็นมี module ข้างนอก มี module ซ้อนเข้าไปอีก ก็คือ eng
 Lenovo@DESKTOP-A96UHVC C:\Users\Lenovo\Desktop\helloapp\eng
$ go mod init exmple.com/helloapp/eng

*/
