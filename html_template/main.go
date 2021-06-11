package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"text/template"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	resp, err := ioutil.ReadFile("C:/Users/Lenovo/Desktop/goworkspace/src/html_template/todo.json")
	//resp, err := http.Get("https://jsonplaceholder.typicode.com/todos")
	if err != nil {
		return
	}
	todoDecoder := json.NewDecoder(bytes.NewReader(resp)) //resp.Body
	todo := []Todo{}
	todoDecoder.Decode(&todo)
	//resp.Body.Close()
	// todoEcoder := json.NewEncoder(os.Stdout)
	// todoEcoder.Encode(todo)

	indexTemplateByte, err := ioutil.ReadFile("C:/Users/Lenovo/Desktop/goworkspace/src/html_template/index.html")
	if err != nil {
		return
	}
	indexTemplate := template.Must(template.New("index").Parse(string(indexTemplateByte)))
	indexTemplate.Execute(os.Stdout, todo)

}
