package main

import (
	"errors"
	"fmt"
)

func withdraw(amount int) (int, error) {
	if amount > 100 {
		return 0, errors.New("insufficient fund")
	}
	return amount, nil
}

func main() {
	amount, err := withdraw(200)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("please collect your money: ", amount)
}

/* Propagate error
คือการส่งต่อ error ให้กับคนที่เรียกมัน สมมติว่ามี error เกิแทนที่เราจะ handle ด้วยตัวเราเอง
เราก็ส่งกลับไปหาผู้เรียก(caller)/caller function

คนถอนเงินออกจากธนาคาร ก็จะมีแอปที่ติดต่อกับดาต้าเบส

go ไม่มีคอนเซปเรื่องการโทรล เราจะรีเทิร์น error กลับไปให้ caller เป็นคนจัดการเอง
ว่าจะเอายังไงดีกับ error ที่เกิดขึ้น

*/
