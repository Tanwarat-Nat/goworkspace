package main

import (
	"fmt"

	gh "github.com/Tanwarat-Nat/weight"
	gl "gitlab.com/Tanwarat-Nat/weight"
)

func main() {
	k := gh.KG(1.5)
	k2 := gl.KG(1.5)
	k3 := gl.KG(1.5)
	fmt.Println(k)
	fmt.Println(gl.KgToLb(k3 + k2))
}

// ต่อให้ code ข้างในเหมือนกัน แต่ import มาจากต่างที่กัน
// ก็จะไม่สามารถทำงานร่วมกันได้
// ในกรณีที่เป็น type declaration เท่านั้น
// ถ้ากับ constant ไม่ปัญหา

// go get sandbox ได้โดยไม่ต้องไป go get ทีละอัน
// จะทำการเข้าไป get ใน sandbox แล้วทำการโหลด dependency ต่างๆ
