package main

import (
	"testing"
)

//func Test[Name] (*testing.T) {}

func TestAddTableDriven(t *testing.T) {
	tests := []struct {
		lhs      int
		rhs      int
		expected int
	}{
		{lhs: 2, rhs: 2, expected: 6},
		{lhs: 1, rhs: 2, expected: 3},
		{lhs: 2, rhs: 2, expected: 4},
		{lhs: 10, rhs: -5, expected: 5},
		{lhs: 10, rhs: -5, expected: 50},
	}
	for _, test := range tests {
		ans := add(test.lhs, test.rhs)
		if ans != test.expected {
			t.Errorf("add(%d,%d) = %d . Wanted %d", test.lhs, test.rhs, ans, test.expected)
		}
	}

}

/*
้ถ้าสมมติว่าเราต้องการจะเทสที่ไม่ใช่ 1,2 ต้องการเทส 2,2 เพระามันเป็น 4 เราก๊อปปี้เทสแบบนั้น
เพื่อที่จะมาเปลี่นพารามิเตอร์แค่ไม่กี่ตัว มันก็ไม่สมเหตุสมผล โค๊ด duplication มากมาย
error ตรงนี้มันก็ซ้ำกันหมด เราสามารถยุบเทส 2 ตัวนั้นมาเป็นอันเดียวได้ ที่เรียกว่า
table driven ที่มันเป็น table เพราะว่า  แทนที่เราจะมาประกาศแต่ละ aggrement ที่จะ
เทสกับฟังก์ชั่นนั้นๆ เป็นคนละฟังก์ชั่น เราก็สามารถเขียนใส่คอลเล็กชั่นของเราได้ นั่นก้คือ tests
เราจะเลือก data structure อะไรก็ได้ แต่ก่อนอื่นเราไปเขียน logic ที่จะใช้ในการเทสก่อน จากนั้น
เราก็เขียน for loop เพิ่มเติม


Errorf มันจะไม่ stop execution มัน fail ใช่มั้ย มันก็ print ออกมา
คนอื่นก็ยังไม่ได้ run ใช่มั้ย ก็ run ต่อ run ต่อกันไป ตรงนี้ fail ก็จะแยกกันไป
การ run ทั้งหมด ข้อดีอย่างนึงคือการ run 1 ครั้งเราสามารถเห็นได้เลยว่าเราผิดพลาดตรงไหนบ้าง
สมมติเราผิดปุ๊บเราเด้งออกมาเลย เราก็จะมีข้อมูลแค่อันเดียว เราก็ไม่รู้ว่า add ของเรา
พอ fix อันเดียวมันจะไป fail อันอื่นอีกรึเปล่า การโชว์ออกมาทั้งหมดครั้งเดียวแบบนี้
ทำให้เราสามารถเห็นปัญหาได้หลายมุมมากขึ้น

และข้อดีของการทำ table test พอมันไม่มี duplication ก็สามารถมาเปลี่ยนแปลง
error message ตกแต่งให้มันสวยงามยังไงก็ได้ ให้เข้าใจและอ่านง่าย เวลาเทสมันเกิดเฟล
จะใส่ index เพิ่มเข้าไปก็ได้ แล้วแต่การดีไซน์

*/
