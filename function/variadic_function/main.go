package main

import "fmt"

func printEachLine(args ...interface{}) {
	for _, v := range args {
		fmt.Println(v)
	}
}

func main() {
	x := []interface{}{"abc", "dfg", "hjk", 123}
	printEachLine(x...)
	printEachLine("abc", "dfg", "hjk", 123)

}

/*
valiadic คือ function ที่รับ aggrement มาแบบ (a ...interface{})
คือตัวแปรที่สามารถเปลี่ยนแปลงค่าได้

ถ้าเราจะประกาศ function นึงที่สามารถรับได้หลายๆ aggrementโดยที่ไม่มี limit
เราก็ใช้ syntex ตรงนั้นได้เหมือนกัน

(args ...string) ตรง...มันจะสร้าง array ขึ้นมาตัวนึงที่ hold พวก
"abc","dfg","hjk" ไว้ แล้วมันก้จะสร้าง slice มาแร็ป array ตัวนี้ args
เราก็จะมองว่ามันเป็น slice ได้
*/
