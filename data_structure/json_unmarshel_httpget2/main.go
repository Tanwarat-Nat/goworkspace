package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Geo struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type Address struct {
	Street  string `json:"street"`
	Suite   string `json:"suite"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
	Geo     Geo    `json:"geo"`
}

type Company struct {
	Name        string `json:"name"`
	CatchPhrase string `json:"catchPhrase"`
	Bs          string `json:"bs"`
}

type Users []struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Address  Address `json:"address"`
	Phone    string  `json:"phone"`
	Website  string  `json:"website"`
	Company  Company `json:"company"`
}

func main() {
	//https://jsonplaceholder.typicode.com/users
	//https://charnnarong.github.io/json-to-go-struct/
	resp, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return
	}
	dataStruct := Users{}
	v := &dataStruct
	json.Unmarshal([]byte(body), v)
	fmt.Println(len(dataStruct))
	dataStruct[0].Name = "Tanwarat"

	// data structure --> std.Output
	result, err := json.MarshalIndent(dataStruct, "", "    ")
	if err != nil {
		return
	}
	fmt.Println(string(result))

}
