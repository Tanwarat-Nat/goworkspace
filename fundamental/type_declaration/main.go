package main

import (
	"fmt"
	"math"
)

// type name underlying-type
type KG float64
type LB float64

func main() {
	k := KG(3)
	b := LB(3)

	// A constant may be given a type explicitly
	// by a constant declaration or conversion
	// คือ type จะถูกระบุโดยชัดเจน
	// อย่าง convert โดยตรง -> float64(3)
	// ในที่นี้ไม่ได้ทำ v
	fmt.Println(k == 3)

	// implicitly when used in a variable declaration
	// or an assignment or as an operand in an expression.
	// คือ type จะถูกเจาะจงโดยตรงแต่จะทำแบบเงียบๆ
	// k + 3 = assignment / k,3 = operand / + = operater
	// error ถ้าไม่สามารถ  represented as a value of the respective type.
	result := k + 3
	result2 := k + math.Pi

	fmt.Println(result)
	fmt.Println(result2)

	fmt.Printf("%T", result)
	fmt.Printf("%T", result2)

	// จะ comparison ได้ type ต้องชนิดเดียวกัน
	// => KG == KG, LB == LB
	// การ class ทำได้เพราะ underlying-type เหมือนกัน
	// b -> KG(b)
	// เป็นการเปลี่ยนแค่ชนิดของข้อมูลเท่านั้นจาก LB -> KG
	fmt.Println(k == KG(b))

}

// การ comparison หรือว่าจะเป็น operation ต่างๆ
// จะขึ้นอยู่กับ underlying-type ว่าสามารถให้ทำอะไรได้บ้าง
// เช่น string ไม่สามารถเอามา หาร แต่เอามาบวก หรือ concastinate กัน

// การ calculation ทำได้โดยการ
// convert untype อย่าง 3 -> KG(3)
// แล้วก็จะเอามา +,-,*,/ กันได้
