package main

import (
	"encoding/json"
	"fmt"
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
	// json.Decode --> data structure
	jsonDecoder := json.NewDecoder(resp.Body)
	dataStruct := Users{}
	jsonDecoder.Decode(&dataStruct)
	resp.Body.Close()
	fmt.Println(dataStruct)
	fmt.Println(len(dataStruct))
	dataStruct[0].Name = "Tanwarat"

}

// decode เรารู้จักว่า json เป็นยังไง แต่คนอื่นไม่รู้ เราต้องทำให้เป็น struct
// json --> struct
// json.Decode --> data structure
