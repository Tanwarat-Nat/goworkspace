package main

import (
	"log"
	"net/http"
	"os"
)

// เราจะสร้าง interface เพื่อทำการ satify listenandserve
// ตรงนี้ก็จะมีอะไรสักอย่างที่ทำการ implement interface ตรงนี้ ซึ่ง

// type Handler interface {
// 	ServeHTTP(ResponseWriter, *Request)
// }

type myHandler func()

//myHandler เป็นไทป์ฟังก์ชั่นที่ไม่รับอะไรเข้ามาเลย แล้ว satify กับ interface

func (m myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted) //set status ได้
	w.Write([]byte("hello"))
}

func main() {
	log.SetFlags(0)
	log.Println(os.Getpid())
	http.ListenAndServe(":8080", myHandler(func() {

	}))
}

/*
หัวใจสำคัญของ ListenAndServe นั่นก็คือตัว second aggrement หรือ parameter ตัวที่ 2
ก็คือ handler จริงแล้วมันก็คือ interface อันนึงที่อยู่ภายใต้ package http

เราจะสร้าง inventory จะมี 2 endpoint หรือจะให้ถูกหน่อยคือมี 2 api
- items
- pricr?item=<name>

ฉะนั้นสิ่งที่เราจะเรียก จะเป็นอะไรแบบนี้
ใน postman
GET --> localhost:8080/items
	--> localhost:8080/pricr?item=apple

*/
