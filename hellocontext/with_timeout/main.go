package main

import (
	"context"
	"log"
	"time"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 500*time.Millisecond)

	ch := time.After(600 * time.Millisecond)

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
อย่างตรงนี้ในเคส 400 * time.Millisecond ch ทำงานเสร็จแล้ว

*/
