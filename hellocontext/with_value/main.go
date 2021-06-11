package main

import (
	"context"
	"fmt"
)

func main() {
	// ทีนี้ไม่ว่าเราจะใส่ไปเท่าไหร่มันก็จะไม่มีทาง error ไม่มีทาง close เราก็สามารถใส่ defer ได้
	ctx, _ := context.WithCancel(context.Background())
	ctx2 := context.WithValue(ctx, "a", "1")
	ctx3 := context.WithValue(ctx2, "b", "2")
	ctx4 := context.WithValue(ctx2, "c", "3")

	lookup(ctx, "ctx", "a")
	lookup(ctx2, "ctx2", "a")
	lookup(ctx3, "ctx3", "a")
	lookup(ctx3, "ctx3", "b")
	lookup(ctx4, "ctx4", "c")

}

func lookup(ctx context.Context, name, k string) {
	fmt.Println(name, ctx.Value(k))
}

/*
ถ้า done มัน close เราก็จะได้ error ถ้า done มันไม่ close error จะเป็น nil
done ไม่ close แสดงว่ามันยังทำงานอยู่ แสดงว่ามันยัง timeout

เวลาเรา cancel ไปเนี่ยมันก็จะ free resource ออกไป
withCancel มันจะ return copy ของ parent มันไม่ได้ return parent
เวลา parent 5^d cancel ลูกๆของมันก็จะถูก cancel ไปด้วย
แต่ context เราสามารถแตกและ wrap ได้หลายๆชั้น อย่างในที่นี้เราโชว์มีแค่ 2 ชั้นเท่านั้น
คืออันแรกเป็น root อันที่ 2 เป็น context ที่เราใช้งานเลยก็ไม่มีปัญหาอะไร




*/
