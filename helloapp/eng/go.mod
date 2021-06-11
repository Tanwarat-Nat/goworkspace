module exmple.com/helloapp/eng

go 1.16

// copy SHA ใน history 06b8da1c8354a8c47856721d1498f160d311ccc1 มาเทียบได้

// Lenovo@DESKTOP-A96UHVC C:\Users\Lenovo\Desktop\helloapp
// $ cd eng

// Lenovo@DESKTOP-A96UHVC C:\Users\Lenovo\Desktop\helloapp\eng
// $ go test -v .
// greet.go:3:8: no required module provides package github.com/Tanwarat-Nat/quote; to add it:
//         go get github.com/Tanwarat-Nat/quote

// Lenovo@DESKTOP-A96UHVC C:\Users\Lenovo\Desktop\helloapp\eng
// $ go get github.com/Tanwarat-Nat/quote
// go: downloading github.com/Tanwarat-Nat/quote v0.0.0-20210525043736-06b8da1c8354
// go get: added github.com/Tanwarat-Nat/quote v0.0.0-20210525043736-06b8da1c8354

// Lenovo@DESKTOP-A96UHVC C:\Users\Lenovo\Desktop\helloapp\eng
// $ go test -v .
// === RUN   TestGreet
// --- PASS: TestGreet (0.00s)
// PASS
// ok      exmple.com/helloapp/eng 0.112s

// check sum (go.sum) มันเป็น cryptography metalrial

require github.com/Tanwarat-Nat/quote v1.0.0
