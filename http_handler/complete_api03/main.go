package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func inventory(w http.ResponseWriter, r *http.Request) {
	// 	r.url /items
	// r.url /price?item=apple
	log.Println("r.url", r.URL.Path)

	switch r.URL.Path {
	case "/items":
		w.Write([]byte("handle items"))
	case "/price": // Path มันจะไม่เอา query param
		w.Write([]byte("handle price"))
	default:
		w.WriteHeader(http.StatusFound)
		fmt.Fprintf(w, "sorry, no such page: %s", r.URL)
	}

}

func main() {
	log.SetFlags(0)
	log.Println(os.Getpid())
	http.ListenAndServe(":8080", http.HandlerFunc(inventory))
}

/*
ต่อไปเราจะมาสร้าง items, price item
localhost:8080/items
ตรง /items มันคือ request url เราจะเก็บข้อมูลตรงนี้ได้
*/
