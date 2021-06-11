package main

import (
	"context"
	"log"
	"time"
)

func main() {
	// ทีนี้ไม่ว่าเราจะใส่ไปเท่าไหร่มันก็จะไม่มีทาง error ไม่มีทาง close เราก็สามารถใส่ defer ได้
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ch := time.After(600 * time.Millisecond)

	// แต่จะเกิดอะไรขึ้นถ้าเราไปเรียก cancel
	time.AfterFunc(400*time.Millisecond, func() {
		log.Println("We are about to cancel thects. ")
		cancel()
	})

	select {
	case <-ctx.Done():
		log.Println("ctx.Done()")
		log.Printf("ctx.Err() Done: %q", ctx.Err())
	case <-ch:
		log.Println("Ch done")
		log.Printf("ctx.Err() Ch  : %q", ctx.Err())
	}

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
