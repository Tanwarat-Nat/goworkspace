package hello

import (
	"exmple.com/helloapp/eng"
	"exmple.com/helloapp/th"
) //เขียน exaple ผิด ตามนั้นนน

//อะไรที่ต่อจาก root gomodule จะต้อง prefix มันด้วย exmple.com/helloapp/

func Hello() string {
	return eng.Greet()
}

func HelloTh() string {
	return th.Greet()
}

/*
 รันเทสใน greet_test.go ผ่านแล้ว แต่พอ cd.. กลับมา รันเทสใน hello_test.go กลับไม่ผ่าน
 เพราะมันพยายามไปดาวน์โหลดตรงนี้ ลิงค์ที่เราพาร์ทเข้าไป exmple.com/helloapp/eng
 อย่างปกติเวลาเราพาร์ท github.com/tanwarat.na/repository ตรงนี้มันก็พยายามไป
 ดาวน์โหลดมาจากอินเตอร์เหมือนกัน แล้วมันไม่เจอมันไม่มี เพระาเรายังไม่ได้ push ไป

 แต่ว่าเรารู้ทั้งรู้แล้วว่า package ตรงนี้มันเป็น local ของเราเรากำลังพัฒนาอยู่ เรายังไม่ได้ push เข้าไป
 เราจะบอกให้ go tool ทราบได้ยังไงว่าตรงนี้มันเป๋็น local ของเรา

 วิธีการ มาที่ go mod เราจะบอก go mod ว่าถ้าเจอ import path ตรงนี้ exmple.com/helloapp/eng
 ให้ไปหาที่ local นะ

ใน go.mod ของ helloapp ให้เพิ่ม 2 อย่างนี้
replace exmple.com/helloapp/eng => ./eng
require exmple.com/helloapp/eng v0.0.0

*/
