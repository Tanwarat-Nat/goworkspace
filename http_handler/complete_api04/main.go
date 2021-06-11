package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type inventory map[string]float64

func (iv inventory) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.URL.Path {
	case "/items":

		for k, v := range iv { //map inventory --> database
			fmt.Fprintf(w, "%s : %.2f\n", k, v)
		}
	case "/price": // Path มันจะไม่เอา query param
		searchItem := r.URL.Query().Get("item") // ["apple","orange"]
		price, ok := iv[searchItem]
		if !ok {
			w.WriteHeader(http.StatusFound)
			fmt.Fprintf(w, "no item: %s", searchItem)
			return
		}
		fmt.Fprintf(w, "%.2f\n", price)
	default:
		w.WriteHeader(http.StatusFound)
		fmt.Fprintf(w, "sorry, no such page: %s", r.URL)
	}

}

func main() {
	log.SetFlags(0)
	log.Println(os.Getpid())
	inven := inventory{
		"apple":  1.25,
		"orange": 0.99,
	}
	http.ListenAndServe(":8080", inven)
}

/*
 เราไม่ได้อยาก return แค่ตัวอักษร เราอยาก return เป็น List of items
 การจะทำอย่างนั้นได้เราต้องมี repository อันนึง ดูเหมือนว่า HandlerFunc ตัวนี้จะไม่พอแล้ว
 เพราะตรงนั้นเป็นแค่ฟังก์ชั่น


*/
