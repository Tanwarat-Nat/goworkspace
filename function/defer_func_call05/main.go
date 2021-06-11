package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("https://golang.org")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}

/*
defer เหมาะจะใช้กับการ close resource (resp.Body.Close())
แทนที่จะเรียกแบบในแต่ละบรรทัดแบบ control flow ปกติ เราสามารถ
เรียกที่ defer ได้ที่เป็น pattern ค่อนข้างยอดฮิต defer resp.Body.Close()
ต่อไปก็ไม่ต้องห่วงว่าโปรแกรมจะจบการทำงาน จะผิดพลาดอะไร defer ก็จะถูก execute

*/
