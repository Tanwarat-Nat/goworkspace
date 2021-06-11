package main

import "net/http"

func main() {
	// run web server on port :8080

	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	}))

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("hello 2"))
	// })

	// http.ListenAndServe(":8080", nil)
}
