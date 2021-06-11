package main

import "fmt"

func main() {
	defer fmt.Println("a")
	defer fmt.Println("b")
	defer fmt.Println("c")
	fmt.Println("hello")
}

/*
defer บอกให้บรรทัดอื่น executed ไปก่อนนะ ให้ทั้งหมด function นี้
execute ไปก่อน สุดท้ายค่อยมา executed บรรทัดนี้ทีหลัง

เฮ้ ทุกคนล่วงหน้าไปก่อนก่อนแลยนะ ไปกันหมดแล้ว ฉันค่อยตามไปทีหลัง ฉันปิดท้ายๆ

defer function call คือ เราจะ defer แต่ยังไม่ call มัน
เราจะไม่ call มันจนกว่าทุกคนทำงานเสร็จ หมายถึงทุกบรรทัดภายใต้
function นี้ทำงานเสร็จ แล้ว function นั้นจะถูกเรียก ตรงนั้นแหละจะเรียกว่า defer

ถ้ามี defer หลายที่ ตัวล่างสุด จะถูก executed ก่อน ตัวเเรกสุด จะถูก executed ทีหลัง
ธรรมชาติของ defer stack คือ first In - last Out
ชิดในอ่ะ เข้าก่อนก็อยู่ใน เข้าทีหลังก็อยู่นอก ตอนออกก็ออกก่อนไง คนในสุดถึงจะออกได้


*/
