package bench

import "math"

// 7867
func isPrime(x int) bool {
	for i := 2; i < x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func isPrimeV2(x int) bool {
	for i := 2; i < int(math.Sqrt(float64(x))); i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func isPrimeV3(x int) bool {
	upper := int(math.Sqrt(float64(x)))
	for i := 2; i < upper; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

/*
เวลาเขียนโค้ดแล้วรู้สึกว่าเขียนโค้ดแบบนี้มันช้าไป ถ้าเปลี่ยนแบบใหม่ ใช้ฟังก์ชั่นนี้มันจะเร็วขึ้น
การมโนพวกนี้เขาเรียกว่า  premature optimization การคิดเอง เออเองว่าทำแบบนี้แล้วมันจะดีขึ้น
จริงๆแล้วมันไม่ใช้ เราต้องมาทดสอบดูว่าเร็วขึ้นจริงรึเปล่า การจะทำแบบนั้นได้เราจะต้องมาทำ
benchmark testing เพื่อพิสูจน์ว่าเเต่ละฟังก์ชั่นที่เราเขียนไปมันช้า เร็วต่างกันขนาดไหน

prime number เช่น 2,3,5,.. เป็นตัวเลขที่มากกว่าหนึ่ง มีแค่หนึ่งกับตัวมันเองที่หารลงตัว ถ้ามันมีตัวใดที่น้อยกว่า
ตัวมันเองแล้วมันคูณกัน แล้วสามารถได้ตัวนั้น ตัวนั้นจะไม่ใช่ prime number เช่น 4,8,9..


*/
