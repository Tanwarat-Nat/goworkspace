package main

import (
	"log"
	"net/http"
	"os"
)

func inventory(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello2"))
}

func main() {
	log.SetFlags(0)
	log.Println(os.Getpid())
	http.ListenAndServe(":8080", http.HandlerFunc(inventory))
}

/*
pattern ด้านบนทุกคนก็ทำกันแบบนี้ ทำกันซ้ำๆ go ก็เลยเห้ย pattern นี้คนทำกันบ่อยมาก
ใช้กันเยอะมาก go ก็เลยสร้างตัว wrapper มาอันนึงซึ่งจะเห็นว่า peattern ด้านบนนั้นมันไม่มีอะไรเลย
myHandler ของเราทำแค่ประกาศ func(http.ResponseWriter, *http.Request) ออกมา
แล้วก็ wrap ตรงนี้ ServeHTTP(w http.ResponseWriter, r *http.Request) เข้าไป
satify interface ตรงนี้ m(w, r) แล้วก็ใช้ logic custom ที่เราเขียนมาเองตรง inventory

*/
