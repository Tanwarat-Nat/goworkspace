package prime

import (
	"fmt"
	"math"
)

// var name type
var notPrimes [1000000]bool //ทั้งหมดไม่ใช่ primes

// ไม่สามารถถูกเรียกใช้โดย program แต่ว่ามันจะถูก involve อัตโนมัติ
// ไม่สามารถเรียกใช้ได้โดยตรง ไม่งั้นมันจะเป็นฟังก์ชั่นธรรมดา
func init() {

	// Algoritnm -> Sieve of Eratosthenes
	fmt.Println("initialization in Prime package")
	sqrtLen := int(math.Ceil(math.Sqrt(float64(len(notPrimes)))))
	for i := 2; i < sqrtLen; i++ {
		if notPrimes[i] {
			continue
		}
		notPrimes[i] = false
		for j := i * 2; j < len(notPrimes); j += i {
			notPrimes[j] = true
		}
	}
	fmt.Println("initialized")
}

func IsPrime(x int) bool {
	return !notPrimes[x]

}

// primes ใน func init จะถูก execute แค่ครั้งเดียว
// และถูก innitialize valiable ต่างๆในระดับ package lavel
// แล้วทุกครั้งเวลาเรามาใช้ function ที่เรา export ไว้ให้ client คนอื่นๆมาใช้
// ก็ทำการ asast local variable ของเรา ในที่นี้คือ notPrimes
