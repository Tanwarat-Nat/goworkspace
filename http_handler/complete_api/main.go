package main

import (
	"log"
	"net/http"
	"os"
)

type myHandler func(http.ResponseWriter, *http.Request)

func (m myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m(w, r) // m มันไทป์ฟังก์ชั่นอยู่แล้วพาร์ทพารามีเตอร์เข้ามาได้เลย
}

func inventory(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func main() {
	log.SetFlags(0)
	log.Println(os.Getpid())
	http.ListenAndServe(":8080", myHandler(inventory))
}

/*
เราจะทำการ refactor เอาฟังก์ชั่นที่ประกาศไว้ตรงนี้ไปไว้ที่อื่น

*/
