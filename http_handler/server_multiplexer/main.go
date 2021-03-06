package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type inventory map[string]float64

func (iv inventory) items(w http.ResponseWriter, r *http.Request) {
	for k, v := range iv {
		fmt.Fprintf(w, "%s : %.2f\n", k, v)
	}
}

func (iv inventory) price(w http.ResponseWriter, r *http.Request) {
	searchItem := r.URL.Query().Get("item")
	price, ok := iv[searchItem]
	if !ok {
		w.WriteHeader(http.StatusFound)
		fmt.Fprintf(w, "no item: %s", searchItem)
		return
	}
	fmt.Fprintf(w, "%.2f\n", price)
}

func (iv inventory) notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusFound)
	fmt.Fprintf(w, "sorry, no such page: %s", r.URL)
}

func main() {
	log.SetFlags(0)
	log.Println(os.Getpid())
	inven := inventory{
		"apple":  1.25,
		"orange": 0.99,
	}

	mux := http.NewServeMux()
	mux.Handle("/items", http.HandlerFunc(inven.items))
	mux.Handle("/price", http.HandlerFunc(inven.price))
	mux.Handle("/", http.HandlerFunc(inven.notFound))

	http.ListenAndServe(":8080", mux)

}

/*
จนตอนนี้เราก็สามารถสร้าง web server ที่สามารถรับ request ได้ทั้งเป็นแบบ item และก็ price
แล้วก็สามารถรับ query param ได้ แล้วก็มาเสิร์ชหาใน inventory ของเรา แต่โค๊ดมันยังยุ่งๆ
เราจะมา introduce ในตัว concepr ใหม่อันนึง คือ server multiplexer

basic cconcept Multiplexer ->
 จริงๆมันมาจากเรื่องของอิเล็คทรอนิกส์ ก็คือการ combiled แต่ละ input ออกมาเป็น output เดียวกัน


*/
