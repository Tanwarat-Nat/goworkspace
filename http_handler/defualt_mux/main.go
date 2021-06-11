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

	http.HandleFunc("/items", inven.items)
	http.HandleFunc("/price", inven.price)
	http.HandleFunc("/", inven.notFound)

	http.ListenAndServe(":8080", nil)

}

/*
อันนี้เราจะมา simplified จาก code ที่เราทำในคลิปที่ผ่านมา

*/
