package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id,omitempty"`
	Title     string `json:"title,omitempty"`
	Completed bool   `json:"completed"`
}

func main() {
	//https://jsonplaceholder.typicode.com/todos
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos")
	if err != nil {
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return
	}
	dataStruct := []Todo{} // แต่ตรงนี้เราประกาศเป็น slice เพระาไม่รู้ size
	v := &dataStruct
	json.Unmarshal([]byte(body), v)
	fmt.Println(len(dataStruct))

	// data structure --> std.Output
	result, err := json.MarshalIndent(dataStruct, "", "    ")
	if err != nil {
		return
	}
	fmt.Println(string(result))

}
