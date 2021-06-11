package main

import (
	"errors"
	"fmt"
	"log"
)

const dbReady = false // บอกให้รู้ว่า db มัน up/down อยู่
const balance = 200

func getBalance() (int, error) {

	if !dbReady {
		return 0, errors.New("getBalance: database is down")
	}
	return balance, nil
}

func withdraw(amount int) (int, error) {

	balance, err := getBalance()
	if err != nil {
		return 0, fmt.Errorf("withdraw: %v", err)
	}

	if amount > balance {
		return 0, errors.New("withdraw: insufficient fund")
	}
	return amount, nil
}

func main() {
	log.SetFlags(0)
	amount, err := withdraw(200)
	if err != nil {
		log.Fatalf("main: %v", err)
		// fmt.Println("main: ", err)
		// os.Exit(1)
	}
	fmt.Println("Please collect your money: ", amount)
}

/*return ต้องบอกได้ด้วยว่า return ได้ success / fail
เพราะถ้ามีโปรแกรมอื่นมาต่อมันจะงง ว่าตกลงมันได้หรือไม่ได้

withdraw เป็น business logic ไม่ควร fail ง่ายๆ

แต่ถ้าเราเขียนโปรแกรมใน unix เช่นพวกใน cp ก็พวกการ copy file unix command
เวลามัน fail มันจะ return status ออกมาเลยว่า fail นะ

principle พวกนี้เรียกว่า fail-fast principle คือ fail-fast system จะหยุดการทำงาน
ของโปรแกรม แทนที่จะพยายาม continune ต่อไป เช่น

ถ้ามัน withdraw ไม่ได้ แต่โปรแกรมมันไม่ได้มีแค่ withdraw นิ มันอาจเอาเงินที่
withdraw มาได้ไป transfer กับ bank ธนบัตรนู้นนี้นั้น

ก็คือถ้าเรา withdraw ไม่ได้ ก็ไม่มีประโยชน์ที่จะมา execute code พวกนี้ เราก็ return กลับไปเลย
แล้วเราก็ต้อง return status เป็น fail นะ

concept การเขียนโปรแกรมใน go เวลา error มันจะ fail-fast รีบ return กลับไปเลย
ความชัดเจนก็จะเห็นได้ง่ายขึ้น อย่างตรงนี้เวลา code ขึ้น production แล้วเรากลับมาอ่านทีหลัง
เรา observe ใน log โอ้ง่ายเลย database is down นะ เราสามารถ dianosys ได้รวดเร็ว

แต่ว่า code ตรงนี้มันเกี่ยวข้องกับ database เราควร refactor มันออกไปเพราะ
มันคงละฟังก์ชั่นกับ withdraw แล้ว
*/
