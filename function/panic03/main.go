package main

import (
	"fmt"
	"os"
)

var cygwinPath = os.Getenv("CYGWINX") //set path

func init() {
	defer func() {
		fmt.Println("clean up in init.")
	}()
	fmt.Println("init cygwin path:", cygwinPath)
	if cygwinPath == "" {
		panic("cannot set cygwin path, make sure you have one")
	}
}

func main() {
	// utilities for Cygwin.
	// path cygwin ?
	fmt.Println("hello to cygwin utilities.")

}

/*utilities for Cygwin
คือ cygwin ที่เราทำอยู่ตรงนี้แหละที่มี libary เยอะๆแล้วเราเอาตัวนี้ไป point ไปจับมัน
แล้วเราต้องการให้แร๊ปเปอรืให้มันง่ายๆ
function ก็อาจจะเรียกในนี้แล้วไปเรียกใน libary อีกที่นึง ดังนั้นโปรแกรมต้องมีอย่างน้อย
ต้องรู้ว่า path ของ cygwin อยู่ที่ไหน

ถ้า set os.Getenv("CYGWINX") -> โปรแกรมพยายามจะอ่านค่ะ
ค่าที่มันเกิดขึ้นจะเป็น emtpy เพราะมันไม่มี ถ้ามัน emtpy แล้วจะมีปย.
อะไรให้โปรแกรมของเราทำงานต่อไป เพราะโปรแกรมของเรามันต้องใช้
มัน depend on cygwinPath
usecase แบบนี้เราพอจะ panic ออกมาได้ ถ้าเป็นแบบนี้เราก็ check เลยว่า if...
แต่ถึงแม้ว่าโปแกรมจะ stop execute flow ทั้งหมด ยะังไงก็ตาม defer ก็ยังจะ run อยู่
แล้วจบด้วยการแสดง panic ออกไป

panic มันจะ bubble มันจะมี function



*/
