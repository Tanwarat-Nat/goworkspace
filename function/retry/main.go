package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

const dbReady = false // บอกให้รู้ว่า db มัน up/down อยู่
const balance = 200
const numberToSuccess = 10

func connectDb(nTry int) error {
	if nTry == numberToSuccess {
		return nil
	}
	return errors.New("busy")
}

func waitForDatabase() error {
	timeout := 3 * time.Second
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		err := connectDb(tries)
		if err == nil {
			return nil
		}
		log.Printf("database is not responding (%v); retrying...", err)
		time.Sleep(time.Second << uint(tries))
		// retry เสร็จแล้วหยุดรอสักพักนึง ไม่ให้โหลดมันมากเกินไป วิธีคือ sleep มัน
		// if err != nil {
		// 	log.Println(err)
		// 	continue // ถ้าเกิด error ขึ้นเราก็ต้อง retry ใหม่ ก็ให้ continune ต่อไป
		// }
		// break // ถ้าไม่ error ก็ออกไป
	}
	return fmt.Errorf("waitForDatabase: timeout %v", timeout)
}

func getBalance() (int, error) {
	//log.Println("retrying...")
	if err := waitForDatabase(); err != nil {
		return 0, fmt.Errorf("getBalance: %v", err)
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

/*application ต่อ database แล้วมัน fail database down เราอาจจะเห้ยย
ขอรองต่ออีกสัก 2 รอบได้มั้ย ว่ามันจะยัง down อยู่รึเปล่าถ้ามากกว่า 3s 5s หรือว่า
ต่อไปเกิน 10 ครั้งแล้วไม่ได้ สุดท้ายแล้วค่อย fail

retry จะเกิดตอน state มันไม่ ready

*/
