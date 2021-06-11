module exmple.com/helloapp

go 1.16

replace exmple.com/helloapp/eng => ./eng // ให้มันไปใช้ที่ local

replace exmple.com/helloapp/th => ./th

require (
	exmple.com/helloapp/eng v0.0.0-00010101000000-000000000000
	exmple.com/helloapp/th v0.0.0-00010101000000-000000000000
)

// Lenovo@DESKTOP-A96UHVC C:\Users\Lenovo\Desktop\helloapp
// $ go get exmple.com/helloapp/eng
// go get: added exmple.com/helloapp/eng v0.0.0-00010101000000-000000000000
