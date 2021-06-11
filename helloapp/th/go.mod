module exmple.com/helloapp/th

go 1.16

require github.com/Tanwarat-Nat/quote/v2 v2.0.0

// check ว่าเราใช้ โมดูล อะไรบ้าง
// Lenovo@DESKTOP-A96UHVC C:\Users\Lenovo\Desktop\helloapp\th
// $ go list -m all
// exmple.com/helloapp/th
//บรรทัดแรกทุกครั้งจะเป็น module name ของมัน เป็น module ที่เรา import path ของมัน
// github.com/Tanwarat-Nat/quote v0.0.0-20210525043736-06b8da1c8354
//บรรทัดถัดไปจะเป็น package ที่มัน depend on มันจะเรียงตามตัวอักษร
//มันไม่ได้มีความหมายว่าอันไหนอยู่บรรทัดที่ 1 จะเป็น root ต่อไปก็จะเป็น parent ไม่ใข่นะ

//Upgrade version
// Lenovo@DESKTOP-A96UHVC C:\Users\Lenovo\Desktop\helloapp\th
// $ go get github.com/Tanwarat-Nat/quote@v0.1.0
// go: downloading github.com/Tanwarat-Nat/quote v0.1.0
// go get: upgraded github.com/Tanwarat-Nat/quote v0.0.0-20210525043736-06b8da1c8354 => v0.1.0
